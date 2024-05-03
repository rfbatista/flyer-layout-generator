from collections.abc import Iterable
from typing import List
from PIL import Image
from pydantic import BaseModel

from app.entities.photoshop import Elemento, PhotoshopElement


class Componente(BaseModel):
    id: int
    elements: List[PhotoshopElement]
    type: str
    width: int
    height: int
    xi: int
    yi: int
    xii: int
    yii: int
    _items: List[Elemento]

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

    def draw_in_image(self, to_image):
        for item in self._items:
            el = [e for e in self.elements if e.layer_id == item.layer_id]
            im = item.image()
            im.thumbnail(el[0].size())
            to_image.paste(im, (self.xi, self.yi), im)
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
        # print(self.elements)
        elements_ids = [e.layer_id for e in self.elements]
        for _, layer in enumerate(layer):
            # print(layer.layer_id)
            if layer.layer_id in elements_ids:
                self._items.append(Elemento(layer))
            self.index_elements(layer, pre + "\t")

    def getImages(self):
        return [x.image() for x in self._items]

    # (left, top, right, bottom)
    def coord(self):
        if len(self._items) == 0:
            return (0, 0)
        bbox = self._items[0].box()
        top_x = bbox[0]
        top_y = bbox[1]
        bot_x = bbox[2]
        bot_y = bbox[3]
        for l in self._items:
            coord = l.box()
            if top_x > coord[0]:
                top_x = coord[0]
            if top_y > coord[1]:
                top_y = coord[1]
            if bot_x < coord[2]:
                bot_x = coord[2]
            if bot_y < coord[3]:
                bot_y = coord[3]
        return (top_x, top_y, bot_x, bot_y)

    def size(self):
        return (self.width, self.height)

    def image(self):
        img = Image.new("RGB", self.size(), color="black")
        coord = self.coord()
        for item in self._items:
            im = item.image()
            img.paste(im, item.position_from((coord[0], coord[1])), im)
        return img

    def draw_in(self, img, point):
        up_left = self.coord()
        move = (up_left[0] - point[0], up_left[1] - point[1])
        for item in self._items:
            im = item.image()
            img.paste(im, item.position_from(move), im)
        return img

    def draw_in_template_position(self, img, template):
        up_left = self.coord()
        move = (up_left[0] - template.xi, up_left[1] - template.yi)
        scale = 1
        if self.width() > self.height():
            if template.width is not None:
                scale = template.width / self.width()
        else:
            if template.height is not None:
                scale = template.height / self.height()
        for item in self._items:
            im = item.image()
            im.thumbnail((int(im.width * scale), int(im.height * scale)))
            img.paste(im, item.position_from(move), im)
        return img
