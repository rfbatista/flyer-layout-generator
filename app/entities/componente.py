from collections.abc import Iterable
from copy import deepcopy
from typing import List, Optional
from pydantic import BaseModel
from PIL import Image

from app.entities.photoshop import Elemento, DesignElement

class Componente(BaseModel):
    id: int
    elements: List[DesignElement]
    type: str
    width: int
    height: int
    xi: int
    yi: int
    xii: int
    yii: int
    bbox_xi: Optional[int] = 0
    bbox_xii:Optional[int] = 0
    bbox_yi: Optional[int] = 0
    bbox_yii:Optional[int] = 0
    _items: List[Elemento] = []


    def bbox_height(self):
        return (self.bbox_yii or 0 ) - (self.bbox_yi or 0 )

    def bbox_width(self):
        return (self.bbox_xii or 0 ) -  (self.bbox_xi or 0 )

    def resize_component_element(
        self, element: DesignElement, width: int, height: int
    ) -> DesignElement:
        nelement = deepcopy(element)
        nelement.xi = int(round(element.xi * (width / element.width)))
        nelement.yi = int(round(element.yi * (height / element.height)))
        nelement.width = width
        nelement.height = height
        nelement.xii = nelement.xi + nelement.width
        nelement.yii = nelement.yi + nelement.height
        return nelement

    def move_to(self, xi, yi):
        x_move = xi - self.xi
        y_move = yi - self.yi
        self.xi = xi
        self.xii = self.width
        self.yi = yi
        self.yii = yi + self.height
        for e in self.elements:
            e.movement(x_move, y_move)

    def resize_component(self, new_width, new_height):
        width_prorp = new_width / self.width
        height_prop = new_height / self.height
        # width_prorp = (self.width * width_proportion) / self.width
        # height_prop = (self.height * height_proportion) / self.height
        self.xi = int(self.xi * width_prorp)
        self.yi = int(self.yi * height_prop)
        self.xii = int(self.xii * width_prorp)
        self.yii = int(self.yii * height_prop)
        self.width = new_width
        self.height = new_height
        nelements = []
        for elem in self.elements:
            nelement = self.resize_component_element(
                elem,
                int(round(elem.width * width_prorp, 0)),
                int(round(elem.height * height_prop, 0)),
            )
            nelements.append(nelement)
        self.elements = nelements

    def bbox(self):
        return ((self.xi, self.yi), (self.xii, self.yii))

    def pos(self):
        return (self.xi, self.yi)

    def set_size(self, new_size):
        self.width = new_size[0]
        self.height = new_size[1]
        self.xii = self.xi + new_size[0]
        self.yii = self.yi + new_size[1]

    def is_in_pixel(self, pixel) -> bool:
        in_x = pixel[0] >= self.xi and pixel[0] <= self.xii
        in_y = pixel[1] >= self.yi and pixel[1] <= self.yii
        if in_x and in_y:
            return True
        else:
            return False

    def draw_in_image(self, to_image, log = False):
        for item in self._items:
            el = [e for e in self.elements if str(e.layer_id) == str(item.layer_id())]
            if len(el) == 0:
                continue
            element = el[0]
            im = item.image()

            if log == True:
                print(element)
            size = element.size()
            im = im.resize((size[0],size[1]), Image.Resampling.LANCZOS)

            if log == True:
                print(im.size)
            # print(
            #     "positioning element in: %s %s with size %s"
            #     % (el[0].layer_id, el[0].pos(), el[0].size())
            # )
            to_image.paste(im, el[0].pos(), im)
        return to_image

    def add_element(self, item):
        self.elements.append(item)

    def print(self, layer, pre=""):
        if not isinstance(layer, Iterable):
            return
        for _, layer in enumerate(layer):
            # print("%s %s %s" % (pre, layer.name, layer.layer_id))
            self.print(layer, pre + "\t")

    def index_elements(self, layer, pre=""):
        if not isinstance(layer, Iterable):
            return
        elements_ids = [e.layer_id for e in self.elements]
        for _, layer in enumerate(layer):
            if str(layer.layer_id) in elements_ids:
                self._items.append(Elemento(layer))
            self.index_elements(layer, pre + "\t")

    def size(self):
        return (self.width, self.height)

    def __str__(self):
        return "component_id %s size: %s position %s" % (self.id, self.size(), self.pos())
    
