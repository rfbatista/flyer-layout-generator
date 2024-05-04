from typing import List, Optional
from uuid import uuid1, uuid4
from sqlalchemy import ForeignKey, String
from sqlalchemy.orm import Mapped, backref, mapped_column, relationship
from pydantic import BaseModel

from app.db import Base
from app.entities.componente import Componente
from app.entities.design_region import DesignTemplateRegion
from app.value_objects.id import create_id


class PositionDto(BaseModel):
    id: int
    xi: int
    width: int
    yi: int
    height: int


class Position(Base):
    __tablename__ = "templates_positions"

    id: Mapped[int] = mapped_column(primary_key=True)
    xi: Mapped[Optional[int]]
    yi: Mapped[Optional[int]]
    width: Mapped[Optional[int]]
    height: Mapped[Optional[int]]

    template_id: Mapped[int] = mapped_column(ForeignKey("templates.id"))
    template = relationship("Template", back_populates="positions")

    def x_center(self) -> int:
        if self.width is None:
            return 0
        if self.xi is None:
            return 0
        return int(round(self.width / 2)) + self.xi

    def y_center(self) -> int:
        if self.height is None:
            return 0
        if self.yi is None:
            return 0
        return int(round(self.height / 2)) + self.yi


class DesginTemplateDistortion(BaseModel):
    x: int
    y: int


class DesignTemplate(BaseModel):
    id: int = 0
    name: str = ""
    width: int
    height: int
    distortion: DesginTemplateDistortion
    background: Optional[Componente] = None

    def regions(self) -> List[DesignTemplateRegion]:
        width = int(self.width / self.distortion.x)
        height = int(self.height / self.distortion.y)
        regions = []
        xi = 0
        yi = 0
        for _ in range(self.distortion.y):
            for _ in range(self.distortion.x):
                r = DesignTemplateRegion(
                    id=create_id(),
                    xi=xi,
                    xii=xi + width,
                    yi=yi,
                    yii=yi + height,
                )
                regions.append(r)
                xi = xi + width
            yi = yi + height
        # print('regioes geradas: ', ["%s %s" % (c.id, c.bbox()) for c in regions])
        return regions

    def set_background(self, c: Componente):
        self.background = c


class Template(Base):
    __tablename__ = "templates"

    id: Mapped[int] = mapped_column(primary_key=True)
    name: Mapped[str] = mapped_column(String(50))
    width: Mapped[Optional[int]]
    height: Mapped[Optional[int]]

    positions = relationship("Position", back_populates="template")

    def size_guide(self):
        if self.width > self.height:
            return self.width
        else:
            return self.height

    @staticmethod
    def from_db(data, positions_) -> DesignTemplate:
        positions = []
        for item in positions_:
            positions.append(
                PositionDto(
                    id=item.id,
                    xi=item.xi,
                    width=item.width,
                    yi=item.yi,
                    height=item.height,
                )
            )
        return DesignTemplate(
            id=data.id,
            name=data.name,
            positions=positions,
            width=data.width,
            height=data.height,
        )
