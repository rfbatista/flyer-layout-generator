#!/usr/local/bin/python
import io
import json
import sys
from typing import Iterable, List
import uuid

import requests
from psd_tools import PSDImage
from pydantic import BaseModel

from app.config import app_config
from app.entities.photoshop import DesignElement, PhotoshopFile
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
            for layer in element:
                # filename = "%s.png" % (uuid.uuid4())
                # filepath = "%s/%s" % (app_config.dist_path, filename)
                img = layer.composite()
                image_url = upload_image(img, layer.name)
                # if img:
                #     img.save(filepath)
                text = ""
                if layer.kind == "type":
                    text = layer.text
                items.append(
                    DesignElement(
                        id=None,
                        photoshop_id=None,
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
                    )
                )
                index_elements(layer, level=level + 1, group_id=layer.layer_id)

        index_elements(psd)
        return ProcessPhotoshopFileResult(
            image_url=design_image_url,
            elements=items, photoshop=photoshopfile, imagepath="", filepath=""
        )
    except Exception as e:
        logger.exception("failed to save photoshop file")
        return {"error": "internal server error \n %s" % (e)}


if __name__ == "__main__":
    # execute only if run as the entry point into the program
    args = sys.argv
    parameters = args[1:]
    result = process_photoshop_file(parameters[0])
    print(json.dumps(result))
