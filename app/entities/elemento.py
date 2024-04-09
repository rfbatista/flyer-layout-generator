class Elemento:
    def __init__(self, layer):
        self.layer = layer

    def box(self):
        return self.layer.bbox

    def image(self):
        im = self.layer.composite()
        return im

    def position_from(self, origin):
        box = self.box()
        return (box[0] - origin[0], box[1] - origin[1])
