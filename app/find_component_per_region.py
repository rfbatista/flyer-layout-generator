from typing import List
from app.entities.component_status import ComponenteStatus
from app.entities.componente import Componente
from app.entities.template import DesignTemplateRegion


def define_components_per_region(
    regions: List[DesignTemplateRegion], components: List[Componente]
) -> List[DesignTemplateRegion]:
    componentes_in = set()
    for reg in regions:
        status_componentes = [ComponenteStatus(c) for c in components]
        for x in range(reg.xi, reg.xii):
            for y in range(reg.yi, reg.yii):
                for comp in status_componentes:
                    if comp.is_in_pixel((x, y)):
                        comp.increase_pixel()
        choosed_component = None
        for comp in status_componentes:
            if choosed_component is None:
                choosed_component = comp
                continue
            elif choosed_component.is_greater(comp):
                choosed_component = comp
        if choosed_component != None:
            componentes_in.add(choosed_component.c.id)
            reg.set_component(choosed_component.c)

    return regions
