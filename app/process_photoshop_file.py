#!/usr/local/bin/python
import io
import json
import sys
from typing import Iterable, List
import uuid

from rembg import remove
import numpy as np
import requests
from psd_tools import PSDImage
from pydantic import BaseModel
from PIL import Image, ImageChops, ImageDraw, ImageOps

from app.config import app_config
from app.entities.photoshop import DesignElement, PhotoshopFile, Property
from app.logger import logger
from app.upload_image import upload_image


class ProcessDesignFileRequest(BaseModel):
    id: int
    filepath: str


class ProcessPhotoshopFileResult(BaseModel):
    image_url: str
    filepath: str
    photoshop: PhotoshopFile
    elements: List[DesignElement]


endpoint_url = "http://localhost:8000/api/v1/design/{}/file"


def find_bounding_box(image):
    # Ensure image has an alpha channel
    if image.mode != "RGBA":
        image = image.convert("RGBA")

    # Split the image into its RGBA components
    r, g, b, a = image.split()

    # Create a mask for non-transparent pixels
    non_transparent_mask = a.point(lambda p: p > 0 and 255)

    # Convert the image to grayscale and create a mask for non-black pixels
    grayscale = image.convert("L")
    non_black_mask = grayscale.point(lambda p: p != 0 and 255)

    # Combine the masks to find non-black, non-transparent pixels
    combined_mask = ImageChops.lighter(non_transparent_mask, non_black_mask)

    # Find the bounding box of the combined mask
    bbox = combined_mask.getbbox()
    print("bbox", bbox)
    return bbox


def autocrop_image(image, border=0):
    # Get the bounding box
    bbox = image.getbbox()
    print("bbox", bbox)

    # Crop the image to the contents of the bounding box
    image = image.crop(bbox)

    # Determine the width and height of the cropped image
    (width, height) = image.size
    print("size", image.size)

    # Add border
    width += border * 2
    height += border * 2

    # Create a new image object for the output image
    cropped_image = Image.new("RGBA", (width, height), (0, 0, 0, 0))

    # Paste the cropped image onto the new image
    cropped_image.paste(image, (border, border))

    # Done!
    return cropped_image, bbox


# https://github.com/crazyhitty/crop-transparent-image/blob/master/crop.py
def readable_resolution(size: tuple):
    return str(size[0]) + "x" + str(size[1])


def is_pixel_alpha(pixel: tuple or int):
    pixel_value = pixel[3] if isinstance(pixel, tuple) else pixel
    return pixel_value == 0


def crop(image):
    # if not is_valid_image_file(image_path):
    #    raise ValueError(image_path, 'is not a valid png image.', 'Only a valid png file is accepted')

    # image = Image.open(image_path, 'r')
    width = image.size[0]
    height = image.size[1]
    pixels = image.load()

    top = 0
    bottom = 0
    left = 0
    right = 0

    for y in range(0, height):
        for x in range(0, width):
            if not is_pixel_alpha(pixels[x, y]):
                if left == 0 or x - 1 < left:
                    left = x - 1
                break

    for y in range(0, height):
        for x in range(0, width):
            if not is_pixel_alpha(pixels[x, y]):
                if top == 0 or y < top:
                    top = y
                break

    for y in reversed(range(0, height)):
        for x in reversed(range(0, width)):
            if not is_pixel_alpha(pixels[x, y]):
                if right == 0 or x + 1 > right:
                    right = x + 1
                break

    for y in reversed(range(0, height)):
        for x in reversed(range(0, width)):
            if not is_pixel_alpha(pixels[x, y]):
                if bottom == 0 or y + 1 > bottom:
                    bottom = y + 1
                break

    if left == -1:
        left = 0

    if top == -1:
        top = 0

    if right == 0:
        right = width

    if bottom == 0:
        bottom = height

    cropped_image = image.crop((left, top, right, bottom))
    # image.close()
    return cropped_image


def bbox(im):
    a = np.array(im)[:, :, :3]  # keep RGB only
    m = np.any(a != [255, 255, 255], axis=2)
    coords = np.argwhere(m)
    y0, x0, y1, x1 = *np.min(coords, axis=0), *np.max(coords, axis=0)
    return (x0, y0, x1 + 1, y1 + 1)


def procImage(img):
    im = img.convert("RGB")
    # Get rid of existing black border by flood-filling with white from top-left corner
    ImageDraw.floodfill(im, xy=(0, 0), value=(255, 255, 255), thresh=10)

    # Get bounding box of text and trim to it
    bbbox = ImageOps.invert(im).getbbox()
    trimmed = im.crop(bbbox)

    # Add new white border, then new black, then new white border
    res = ImageOps.expand(trimmed, border=10, fill=(255, 255, 255))
    res = ImageOps.expand(res, border=5, fill=(0, 0, 0))
    res = ImageOps.expand(res, border=5, fill=(255, 255, 255))
    return res


alpha_threshold = 100  # Adjust the threshold as needed


def remove_bg(img):
    datas = img.getdata()
    new_data = []
    for item in datas:
        # Change all pixels with alpha greater than the threshold to transparent
        if item[3] < alpha_threshold:
            new_data.append((255, 255, 255, 0))  # Set to fully transparent
        else:
            new_data.append(item)
    img.putdata(new_data)
    return img


def process_photoshop_file(req: ProcessDesignFileRequest):
    try:
        filepath = req.filepath
        res = requests.get(endpoint_url.format(req.id))
        content = res.content
        psd = PSDImage.open(io.BytesIO(content))
        filename = "%s" % (uuid.uuid4())
        img = psd.composite()
        design_image_url = upload_image(img, filename)
        photoshopfile = PhotoshopFile(
            filename="",
            image_path=filename,
            image_extension="png",
            filepath=filepath,
            width=psd.width,
            height=psd.height,
        )

        items = []

        def index_elements(element: PSDImage, level=0, group_id=0):
            if not isinstance(element, Iterable):
                return
            level = 0
            for layer in element:
                # filename = "%s.png" % (uuid.uuid4())
                # filepath = "%s/%s" % (app_config.dist_path, filename)
                img1 = layer.composite()
                image_url = upload_image(img1, layer.name)
                img, img_bbox = autocrop_image(remove_bg(img1))
                # if img:
                #     img.save(filepath)
                text = ""
                if layer.kind == "type":
                    text = layer.text
                print("\n ", layer.name)
                print("before", img_bbox)
                print("after", img.getbbox())
                box = (layer.left, layer.top, layer.right, layer.bottom)
                if img_bbox is not None:
                    box = (
                        layer.left + img_bbox[0],
                        layer.top + img_bbox[1],
                        layer.left + img_bbox[2],
                        layer.top + img_bbox[3],
                    )
                inner_xi = box[0]
                inner_xii = box[2]
                inner_yi = box[1]
                inner_yii = box[3]
                if box[2] > psd.width:
                    inner_xii = psd.width
                if box[3] > psd.height:
                    inner_yii = psd.height
                if box[0] < 0:
                    inner_xi = 0
                if box[1] < 0:
                    inner_yi = 0
                properties = []
                try:
                    if layer.kind == "type":
                        properties = extract_properties(layer)
                except:
                    print("error parsing photoshop file")

                items.append(
                    DesignElement(
                        id=None,
                        photoshop_id=None,
                        inner_xi=inner_xi,
                        inner_xii=inner_xii,
                        inner_yi=inner_yi,
                        inner_yii=inner_yii,
                        xi=layer.left,
                        yi=layer.top,
                        xii=layer.right,
                        yii=layer.bottom,
                        image_extension="png",
                        width=layer.width,
                        height=layer.height,
                        kind=layer.kind,
                        name=layer.name,
                        text=text,
                        is_group=layer.is_group(),
                        group_id=group_id,
                        layer_id=str(layer.layer_id),
                        level=level,
                        image=image_url,
                        properties=properties,
                    )
                )
                level += 1
                index_elements(layer, level=level + 1, group_id=layer.layer_id)

        index_elements(psd)
        return ProcessPhotoshopFileResult(
            image_url=design_image_url,
            elements=items,
            photoshop=photoshopfile,
            filepath="",
        )
    except Exception as e:
        logger.exception("failed to save photoshop file")
        return {"error": "internal server error \n %s" % (e)}

def extract_properties(layer):
    properties = []
    properties.append(Property(key="text", value=layer.text))
    if "FontSet" in layer.resource_dict:
        for font in layer.resource_dict["FontSet"]:
            # print(font)
            # print(font["Name"])
            properties.append(
                Property(key="font_name", value=str(font["Name"]))
            )
    if "StyleSheetSet" in layer.resource_dict:
        for style in layer.resource_dict["StyleSheetSet"]:
            # print(font)
            # print(font["Name"])
            if "StyleSheetData" in style and "FontSize" in style["StyleSheetData"]:
                properties.append(
                    Property(key="font_size", value=str(style["StyleSheetData"]["FontSize"]))
                )
    if "StyleRun" in layer.resource_dict:
        for st in layer.resource_dict["StyleRun"]:
            if (
                "StyleSheet" in st
                and "StyleSheetData" in st["StyleSheet"]
                and "FontSize" in st["StyleSheet"]["StyleSheetData"]
            ):
                properties.append(
                    Property(
                        key="font_size",
                        value=str(
                            st["StyleSheet"]["StyleSheetData"][
                                "FontSize"
                            ]
                        ),
                    )
                )
    return properties

if __name__ == "__main__":
    # execute only if run as the entry point into the program
    args = sys.argv
    parameters = args[1:]
    result = process_photoshop_file(parameters[0])
    print(json.dumps(result))
