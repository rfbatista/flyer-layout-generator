from typing import List

from copy import deepcopy

from pydantic import BaseModel
from app.entities.componente import Componente
from app.entities.photoshop import DesignElement
from app.entities.prancheta import DesignPrancheta


class Dimension(BaseModel):
    width: int
    height: int


def resize_component_element(
    element: DesignElement, width: int, height: int
) -> DesignElement:
    nelement = deepcopy(element)
    nelement.width = width
    nelement.height = height
    nelement.xi = int(round(element.xi * (width / element.width)))
    nelement.yi = int(round(element.yi * (height / element.height)))
    nelement.xii = nelement.xi + nelement.width
    nelement.yii = nelement.yi + nelement.height
    return nelement


def resize_component(
    component: Componente, width_proportion: float, height_proportion: float
) -> Componente:
    width_prorp = (component.width * width_proportion) / component.width
    height_prop = (component.height * height_proportion) / component.height
    ncomponent = deepcopy(component)
    ncomponent.xi = int(component.xi * width_prorp)
    ncomponent.yi = int(component.yi * height_prop)
    ncomponent.xii = int(component.xii * width_prorp)
    ncomponent.yii = int(component.yii * height_prop)
    nelements = []
    for elem in component.elements:
        nelement = resize_component_element(
            elem,
            int(round(elem.width * width_prorp, 0)),
            int(round(elem.height * height_prop, 0)),
        )
        nelements.append(nelement)
    ncomponent.elements = nelements
    return ncomponent


def resize_components(c: List[Componente], width_proportion: float, height_proportion: float) -> List[Componente]:
    ncomponents = []
    for comp in c:
        ncomponent = resize_component(comp, width_proportion, height_proportion)
        ncomponents.append(ncomponent)
    return ncomponents


def distort_image_to(prancheta: DesignPrancheta, to: Dimension) -> DesignPrancheta:
    width_prorp = to.width / prancheta.width
    height_prop = to.height / prancheta.height
    prancheta.components = resize_components(prancheta.components, width_prorp, height_prop)
    return prancheta
