from collections.abc import Iterable

from app.entities.elemento import Elemento


class Componente:
    def __init__(self, items) -> None:
        self.items = items
        self._items = []

    def print(self, layer, pre=""):
        if not isinstance(layer, Iterable):
            return
        for idx, layer in enumerate(layer):
            print("%s %s %s" % (pre, layer.name, layer.layer_id))
            self.print(layer, pre + "\t")

    def index_elements(self, layer, pre=""):
        if not isinstance(layer, Iterable):
            return
        for idx, layer in enumerate(layer):
            if layer.layer_id in self.items:
                self._items.append(Elemento(layer))
            self.index_elements(layer, pre + "\t")

    def getImages(self):
        return [x.composite() for x in self._items]

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
        coord = self.coord()
        return (coord[2] - coord[0], coord[3] - coord[1])

    def width(self):
        return self.size()[0]

    def height(self):
        return self.size()[1]

    def image(self):
        img = Image.new("RGB", self.size(), color="black")
        coord = self.coord()
        for item in self._items:
            im = item.image()
            img.paste(im, item.position_from((coord[0], coord[1])), im)
        return img

    def draw(self, crop=False):
        img = self.image()
        if crop:
            img = img.crop(crop)
        plt.imshow(img)
        plt.axis("off")  # Hide axis
        plt.show()

    def draw_in(self, img, point):
        up_left = self.coord()
        move = (up_left[0] - point[0], up_left[1] - point[1])
        for item in self._items:
            im = item.image()
            img.paste(im, item.position_from(move), im)
        return img
