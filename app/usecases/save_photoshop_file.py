from typing import Iterable
import uuid

from fastapi import HTTPException, UploadFile
from psd_tools import PSDImage
from sqlalchemy.orm.session import Session

from app.logger import logger
from app.config import app_config
from app.entities.photoshop import PhotoshopElement, PhotoshopFile


def save_photoshop_file(file: UploadFile, db: Session):
    try:
        unique_id = uuid.uuid4()
        path = "%s/%s" % (app_config.photoshop_files_path, unique_id)
        with open(path, "wb") as f:
            f.write(file.file.read())
        psd = PSDImage.open(path)
        photoshopfile = PhotoshopFile(
            filename=file.filename,
            filepath=path,
            width=psd.width,
            height=psd.height,
        )
        db.add(photoshopfile)
        db.commit()
        db.flush()
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
                items.append(
                    PhotoshopElement(
                        xi=layer.left,
                        yi=layer.top,
                        xii=layer.right,
                        kind=layer.kind,
                        name=layer.name,
                        yii=layer.bottom,
                        is_group=layer.is_group(),
                        group_id=group_id,
                        layer_id=layer.layer_id,
                        level=level,
                        photoshop_id=photoshopfile.id,
                        image=filename,
                    )
                )
                index_elements(layer, level=level + 1, group_id=layer.layer_id)
        index_elements(psd)
        db.add_all(items)
        db.commit()
        db.flush()
        return {
            "file": file.filename,
            "content": file.content_type,
            "path": path,
        }
    except Exception as e:
        logger.exception("failed to save photoshop file")
        raise HTTPException(status_code=500, detail="internal server error \n %s" % (e))
