import datetime
from typing import Iterable, Optional

from pydantic import BaseModel
from PIL.Image import Image
from psd_tools import PSDImage
from sqlalchemy import Column, DateTime, ForeignKey, Integer, String, Boolean, func
from sqlalchemy.orm import Mapped, mapped_column, relationship

from app.db import Base


class Elemento:
    def __init__(self, layer):
        self.layer = layer

    def box(self):
        return self.layer.bbox

    def image(self) -> Image:
        im = self.layer.composite()
        return im

    def position_from(self, origin):
        box = self.box()
        return (box[0] - origin[0], box[1] - origin[1])


class PhotoshopElement(BaseModel):
    xi: int
    kind: str
    text: str
    name: str
    yi: int
    xii: int
    yii: int
    width: int
    height: int
    level: int
    group_id: int
    layer_id: str
    component_color: str | None = None
    is_group: bool
    is_background: bool = False
    image: str

    def __repr__(self):
        return "PhotoshopElement(" + self.image + ")"

    def is_component(self) -> bool:
        return False


class PhotoshopFile(BaseModel):
    filename: str
    filepath: str
    width: int
    height: int
