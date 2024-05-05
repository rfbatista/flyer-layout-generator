from typing import Optional

from pydantic import BaseModel
from PIL.Image import Image


class Elemento:
    def __init__(self, layer):
        self.layer = layer

    def box(self):
        return self.layer.bbox

    def layer_id(self):
        return self.layer.layer_id

    def image(self) -> Image:
        im = self.layer.composite(force=True)
        return im

    def position_from(self, origin):
        box = self.box()
        return (box[0] - origin[0], box[1] - origin[1])

    def size(self):
        return (self.layer.width, self.layer.height)

    def pos(self):
        return (self.box()[0][0], self.box()[0][1])

    def __str__(self):
        return "element_id %s size: %s position %s" % (
            self.layer_id(),
            self.size(),
            self.pos(),
        )


class DesignElement(BaseModel):
    id: Optional[int] = None
    xi: int
    xii: int
    yi: int
    yii: int
    layer_id: str
    width: int
    height: int
    kind: str
    name: Optional[str] = None
    is_group: bool
    group_id: int
    level: int
    photoshop_id: Optional[int] = None
    text: Optional[str] = None
    component_id: Optional[int] = None
    image: Optional[str] = None
    image_extension: Optional[str] = None

    def __repr__(self):
        return 'PhotoshopElement("%s")' % (self.image)

    def is_component(self) -> bool:
        return False

    def size(self):
        return (self.width, self.height)

    def pos(self):
        return (self.xi, self.yi)

    def movement(self, xi_mov, yi_mov):
        self.xi = self.xi + xi_mov
        self.yi = self.yi + yi_mov
        self.xii = self.xi + self.width
        self.yii = self.yi + self.height

    def __str__(self):
        return "design_element_id %s name %s size: %s position %s" % (
            self.id,
            self.name,
            self.size(),
            self.pos(),
        )


class PhotoshopFile(BaseModel):
    filename: str = ""
    filepath: str = ""
    file_extension: Optional[str] = None
    image_path: str = ""
    image_extension: Optional[str] = None
    width: int
    height: int
    id: Optional[int] = None
