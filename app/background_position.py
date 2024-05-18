from app.entities.componente import Componente
from app.entities.template import DesignTemplate

def resize_rectangle_to_fit(input_width, input_height, target_width, target_height, comp: Componente):
    width_ratio = target_width / input_width
    height_ratio = target_height / input_height
    proportion = max(width_ratio, height_ratio)
    new_size = (comp.width * proportion, comp.height * proportion)
    print("bbox", input_width, input_height,"proportion choose", proportion, "prooportions", width_ratio, height_ratio,"target", target_width, target_height, "new_size", new_size)
    return new_size

def resize_background(background: Componente, prancheta: DesignTemplate):
    # new_size = fit_in_region(prancheta.width, prancheta.height, background)
    # new_size = resize_rectangle_to_fit(background.width,background.height,prancheta.width,prancheta.height)
    new_size = resize_rectangle_to_fit(background.bbox_width(),background.bbox_height(),prancheta.width,prancheta.height, background)
    print(prancheta.width, prancheta.height)
    print(new_size[0], new_size[1])
    background.resize_component(new_size[0], new_size[1])
    return background
    
