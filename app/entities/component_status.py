from app.entities.componente import Componente


class ComponenteStatus:
    def __init__(self, comp: Componente):
        self.filling_percentage = 0
        self.reading_priority = 0
        self.pixels = 0
        self.c = comp
        self.comp_type = comp.type

    def is_in_pixel(self, pixel):
        if self.c.is_in_pixel(pixel):
            self.increase_pixel()

    def increase_pixel(self):
        self.pixels += 1

    def is_greater(self, cs) -> bool:
        if cs.pixels <= self.pixels:
            return True
        else:
            return False

    def __repr__(self):
        return (
            str(self.c.id)
            + " "
            + str(self.reading_priority)
            + " "
            + str(self.pixels)
        )
