import datetime
from typing import Iterable, Optional

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


class PhotoshopElement(Base):
    __tablename__ = "photoshop_elements"

    id = Column(Integer, primary_key=True, index=True)
    xi: Mapped[Optional[int]]
    kind: Mapped[Optional[str]]
    name: Mapped[Optional[str]]
    yi: Mapped[Optional[int]]
    xii: Mapped[Optional[int]]
    yii: Mapped[Optional[int]]
    level: Mapped[Optional[int]]
    group_id: Mapped[Optional[int]]
    layer_id: Mapped[Optional[int]]
    component_id = mapped_column(String(46))
    component_color = mapped_column(String(7))
    is_group = mapped_column(Boolean(), default=False)
    is_background = mapped_column(Boolean(), default=False)
    image = mapped_column(String(46))

    photoshop_id: Mapped[int] = mapped_column(ForeignKey("photoshop_files.id"))
    photoshopfile = relationship("PhotoshopFile", foreign_keys=[photoshop_id])

    def __repr__(self):
        return "PhotoshopElement(" + self.image + ")"

    def is_component(self) -> bool:
        return self.component_id is not None


class PhotoshopFile(Base):
    __tablename__ = "photoshop_files"

    id = Column(Integer, primary_key=True, index=True)
    filename = Column(String(100), nullable=False)
    filepath = Column(String(255), nullable=False)
    width = Column(Integer)
    height = Column(Integer)

    elements = relationship(
        "PhotoshopElement",
        primaryjoin="and_(PhotoshopFile.id==PhotoshopElement.photoshop_id)",
    )


