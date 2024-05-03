from typing import List
from app.entities.componente import Componente
from app.entities.template import DesignTemplateRegion

PAD_SIZE = 10


def find_position_to_center(regiao: DesignTemplateRegion, comp: Componente):
    element_size = (comp.width, comp.height)
    bbox = regiao.bbox()
    regiao_center_x = ((bbox[1][0] - bbox[0][0]) / 2) + bbox[0][0]
    regiao_center_y = ((bbox[1][1] - bbox[0][1]) / 2) + bbox[0][1]
    center_position = (
        regiao_center_x - (element_size[0] / 2),
        regiao_center_y - (element_size[1] / 2),
    )
    return (int(center_position[0]), int(center_position[1]))


def add_padding(padding=0, size=(0, 0)):
    new_size = (size[0] - padding, size[1] - padding)
    return new_size


def position_components_in_regions(
    regions: List[DesignTemplateRegion],
) -> List[DesignTemplateRegion]:
    for region in regions:
        if region.component is None:
            continue
        pos = find_position_to_center(region, region.component)
        region.component.xi = pos[0]
        region.component.yi = pos[1]
        new_size = add_padding(PAD_SIZE, region.component.size())
        region.component.set_size(new_size)
    return regions
