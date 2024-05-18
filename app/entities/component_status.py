from app.entities.componente import Componente

type_value = {
    'background': 0,
    'logotipo_marca': 10,
    'logotipo_produto': 9,
    'packshot': 2,
    'celebridade': 3,
    'modelo': 4,
    'ilustracao': 5,
    'oferta': 5,
    'texto_legal': 1,
    'grafismo': 1,
    'texto_cta': 9,
}

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
    
    def comp_indice(self):
        return type_value.get(self.comp_type) or 0


    def __repr__(self):
        return (
            str(self.c.id)
            + " "
            + str(self.reading_priority)
            + " "
            + str(self.pixels)
        )
