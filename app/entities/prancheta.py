import random
from typing import List

from PIL import ImageDraw, Image
from app.entities.componente import Componente
from app.entities.guia import Guia
from app.entities.template import Template


class Prancheta:
    def __init__(self, width, height) -> None:
        self.width = width
        self.height = height
        self.lines = []
        self.componentes: List[Componente] = []
        self.points = []

    def size(self):
        return (self.width, self.height)

    def _line_intersection(self, line1, line2):
        xdiff = (line1[0][0] - line1[1][0], line2[0][0] - line2[1][0])
        ydiff = (line1[0][1] - line1[1][1], line2[0][1] - line2[1][1])

        def det(a, b):
            return a[0] * b[1] - a[1] * b[0]

        div = det(xdiff, ydiff)
        if div == 0:
            return None
        d = (det(*line1), det(*line2))
        x = det(d, xdiff) / div
        y = det(d, ydiff) / div
        return (int(x), int(y))

    def _calc_position_points(self):
        for idx, line_ref in enumerate(self.lines):
            for line in self.lines[idx:]:
                point = self._line_intersection(line_ref, line)
                if point:
                    self.points.append(point)

    def set_pivot_x(self, componente: Componente):
        self.pivot_x = componente

    def set_pivot_y(self, componente: Componente):
        self.pivot_y = componente

    def set_prancheta(self, width: int, height: int):
        self.prancheta_height = height
        self.prancheta_width = width

    def _hide_outer_points(self):
        max_x = max(self.points, key=lambda x: x[0])
        max_y = max(self.points, key=lambda x: x[1])
        self.points = [tup for tup in self.points if tup[0] != max_x[0]]
        self.points = [tup for tup in self.points if tup[1] != max_y[1]]

    def calculate(self, hide_outer_points=True):
        guia = Guia()
        guia.set_prancheta(self.width, self.height)
        guia.set_pivot(self.pivot_x.width(), self.pivot_y.height())
        guia.calculate()
        [self.add_line(line) for line in guia.lines]
        self._calc_position_points()
        if hide_outer_points:
            self._hide_outer_points()

    def add_line(self, line):
        self.lines.append(line)

    def add_componente(self, comp: Componente):
        self.componentes.append(comp)

    def set_guia(self, guia: Guia):
        self.guia = guia

    def set_background(self, comp, crop=None):
        self.background = comp
        self.background_crop = crop

    def image(self):
        img = None
        if self.background:
            img = self.background.image()
            if self.background_crop:
                img = img.crop(self.background_crop)
            img = img.crop((0, 0, self.width, self.height))
        else:
            img = Image.new("RGB", self.size())
        draw = ImageDraw.Draw(img)
        point = random.choice(self.points)
        comp = self.pivot_x
        comp.draw_in(img, point)

        point = random.choice(self.points)
        comp = self.pivot_y
        comp.draw_in(img, point)
        return img

    def image_from_template(self, template: Template) -> Image.Image:
        img = None
        if self.background:
            img = self.background.image()
            if img.width > img.height:
                img.thumbnail((int((img.height*template.height)/template.width), template.height))
            else:
                img.thumbnail((template.width, int((img.width*template.width)/template.height)))
            left = (img.width - template.width) // 2
            upper = (img.height - template.height) // 2
            right = left + template.width
            lower = upper + template.height
            
            # Crop the image based on the calculated coordinates
            pos = (left, upper, right, lower)
            print(pos)
            img = img.crop(pos)
            # if self.background_crop:
            #     img = img.crop(self.background_crop)
            # img = img.crop((0, 0, self.width, self.height))
        else:
            img = Image.new("RGB", self.size())
        img.resize((template.width, template.height))
        components = self.componentes
        for pos in template.positions:
            if len(components) == 0:
                return img
            comp = random.choice(components)
            components.remove(comp)
            comp.draw_in_template_position(img, pos)
        return img
