from typing import List, Optional
from pydantic import BaseModel

from app.entities.componente import Componente
from app.entities.design_region import DesignTemplateRegion
from app.value_objects.id import create_id


class PositionDto(BaseModel):
    id: int
    xi: int
    width: int
    yi: int
    height: int


class Position(BaseModel):
    id: int = 0
    xi: Optional[int] = 0
    yi: Optional[int] = 0
    width: Optional[int] = 0
    height: Optional[int] = 0

    template_id: int = 0

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
            xi = 0
        # print('regioes geradas: ', ["%s %s" % (c.id, c.bbox()) for c in regions])
        return regions

    def set_background(self, c: Componente):
        self.background = c


class Template:
    id: int
    name: str
    width: int
    height: int
    positions = []

    def size_guide(self):
        if self.width > self.height:
            return self.width
        else:
            return self.height

