#!/usr/local/bin/python
import json
import sys
from typing import Iterable, List
import uuid

from psd_tools import PSDImage
from pydantic import BaseModel

from app.config import app_config
from app.entities.photoshop import PhotoshopElement, PhotoshopFile
from app.logger import logger


class ProcessPhotoshopFileRequest(BaseModel):
    filepath: str


class ProcessPhotoshopFileResult(BaseModel):
    imagepath: str
    filepath: str
    photoshop: PhotoshopFile
    elements: List[PhotoshopElement]


def process_photoshop_file(filepath: str):
    try:
        psd = PSDImage.open(filepath)
        filename = "%s" % (uuid.uuid4())
        filepath_save = "%s/%s.png" % (app_config.dist_path, filename)
        img = psd.composite()
        if img:
            img.save(filepath_save)
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
            for layer in element:
                filename = "%s.png" % (uuid.uuid4())
                filepath = "%s/%s" % (app_config.dist_path, filename)
                img = layer.composite()
                if img:
                    img.save(filepath)
                text = ""
                if layer.kind == "type":
                    text = layer.text

                items.append(
                    PhotoshopElement(
                        id=None,
                        photoshop_id=None,
                        xi=layer.left,
                        image_extension="png",
                        yi=layer.top,
                        xii=layer.right,
                        width=layer.width,
                        height=layer.height,
                        kind=layer.kind,
                        name=layer.name,
                        text=text,
                        yii=layer.bottom,
                        is_group=layer.is_group(),
                        group_id=group_id,
                        layer_id=str(layer.layer_id),
                        level=level,
                        image=filepath,
                    )
                )
                index_elements(layer, level=level + 1, group_id=layer.layer_id)

        index_elements(psd)
        return {
            "photoshop": photoshopfile,
            "elements": items,
        }
    except Exception as e:
        logger.exception("failed to save photoshop file")
        return {"error": "internal server error \n %s" % (e)}


if __name__ == "__main__":
    # execute only if run as the entry point into the program
    args = sys.argv
    parameters = args[1:]
    result = process_photoshop_file(parameters[0])
    print(json.dumps(result))
