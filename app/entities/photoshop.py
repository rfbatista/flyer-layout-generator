from typing import Optional

from pydantic import BaseModel
from PIL.Image import Image


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
    id: Optional[str] = None
    xi: int
    xii: int
    yi: int
    yii: int
    layer_id: str
    width: int
    height: int
    kind: str
    name: str
    is_group: bool
    group_id: int
    level: int
    photoshop_id: Optional[int] = None
    text: str
    component_id: Optional[int] = None
    image: Optional[str] = None
    image_extension: Optional[str] = None

    def __repr__(self):
        return 'PhotoshopElement("%s")' % (self.image)

    def is_component(self) -> bool:
        return False


class PhotoshopFile(BaseModel):
    filename: str = ''
    filepath: str = ''
    file_extension: Optional[str] = None
    image_path: str = ''
    image_extension: Optional[str] = None
    width: int
    height: int
