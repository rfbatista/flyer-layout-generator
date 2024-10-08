from typing import List
from app.entities.componente import Componente
from app.entities.template import DesignTemplateRegion


def fit_in_region(region: DesignTemplateRegion, comp: Componente, PAD_SIZE = 10):
    # Calculate the scaling factor
    scale_factor = 1
    region_scale_factor = 1
    if region.width() > region.height():
        region_scale_factor = region.height()
    else:
        region_scale_factor = region.width()
    if comp.width > comp.height:
        scale_factor = (region_scale_factor - PAD_SIZE) / comp.width
    else:
        scale_factor = (region_scale_factor - PAD_SIZE) / comp.height

    # Calculate the new size of the square
    new_size = tuple(int(dim * scale_factor) for dim in comp.size())

    return new_size


def find_position_to_center(regiao: DesignTemplateRegion, comp: Componente, log = False):
    element_size = (comp.width, comp.height)

    if log == True:
        print("element size ", element_size)

    bbox = regiao.bbox()
    if log == True:
        print("bbox ", bbox)
    regiao_center_x = bbox[0][0] + (((bbox[1][0] - bbox[0][0]) - element_size[0]) // 2)
    regiao_center_y = bbox[0][1] + (((bbox[1][1] - bbox[0][1]) - element_size[1]) // 2)
    center_position = (
        regiao_center_x,
        regiao_center_y,
    )
    if log == True:
        print("center position ", center_position)
    return (int(center_position[0]), int(center_position[1]))


def add_padding(padding=0, size=(0, 0)):
    new_size = (size[0] - padding, size[1] - padding)
    return new_size


def position_components_in_regions(
    regions: List[DesignTemplateRegion], log = False
) -> List[DesignTemplateRegion]:
    for region in regions:
        if region.component is None:
            continue
        pad_size = 10
        print(region.component.type)
        if region.component.type == "logotipo_marca" or region.component.type == "logotipo_produto":
            pad_size = 10
        if region.component.type == "modelo":
            pad_size = 0
        if region.component.type == "logotipo_produto":
            pad_size = 0
        new_size = fit_in_region(region, region.component, pad_size)
        if log == True:
            print(new_size)
        region.component.resize_component(new_size[0], new_size[1])
        pos = find_position_to_center(region, region.component)
        # print("region: %s size %s" % (region.id, new_size))
        region.component.move_to(pos[0], pos[1])
        # new_size = add_padding(PAD_SIZE, region.component.size())
        # region.component.set_size(new_size)
    return regions
