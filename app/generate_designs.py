from copy import deepcopy
from datetime import datetime, timezone


from app.background_position import resize_background
from app.distortion import Dimension, distort_image_to
from app.drawer import renderer
from app.entities.prancheta import DesignPrancheta
from app.find_component_per_region import define_components_per_region
from app.position_components import position_components_in_regions
from app.requests import GenerateDesignRequest, GenerateDesignResult


def generate_design(req: GenerateDesignRequest, log = False):
    started_at = datetime.now(timezone.utc)
    background = None
    bg_list = [c for c in req.components if c.type == "background"]
    if len(bg_list) > 0:
        background = bg_list[0]
    req.components = [c for c in req.components if c.type != "background"]
    prancheta = DesignPrancheta(
        template=req.template,
        components=list(filter(lambda c: c.width != 0, req.components)),
        width=req.photoshop.width,
        height=req.photoshop.height,
    )
    to_distort_prancheta = deepcopy(prancheta)
    distorted_prancheta = distort_image_to(
        to_distort_prancheta, Dimension(width=req.template.width, height=req.template.height)
    )
    regions = req.template.regions()
    
    if log == True:
        print("\nregioes definidas")
        for r in regions:
            print(r)


    regions_with_components = define_components_per_region(
        regions, distorted_prancheta.components
    )
    for r in regions_with_components:
        ncomp = None
        for c in req.components:
            if c is not None and r.component is not None:
                if c.id == r.component.id:
                    ncomp = c
        if ncomp is not None:
            r.set_component(ncomp)

    regions_with_components = [r for r in regions_with_components if r.component is not None]
    if log == True:
        print("\nregions with component")

    if log == True:
        for c in regions_with_components:
            print(c.component)
            if c.component is not None:
                for e in c.component.elements:
                    print(e)


    regions_with_components_positioned = position_components_in_regions(
        regions_with_components
    )

    if log == True:
        print("\nregions with component positioned")

    if log == True:
        for c in regions_with_components_positioned:
            print(c.component)
            if c.component is not None:
                for e in c.component.elements:
                    print(e)

    componentes = [
        c.component
        for c in regions_with_components_positioned
        if c.component is not None
    ]

    if log == True:
        print("\nselected components")
        for c in componentes:
            print(c)
            for e in c.elements:
                print(e)

    if background is not None:
        resized_bg = resize_background(background, req.template)
        background = resized_bg

    # if log == True:
    # print("background", background.width, background.height, "template", req.template.width, req.template.height)
    image = renderer.render_png(componentes, req, background)
    finished_at = datetime.now(timezone.utc)
    return GenerateDesignResult(
        photoshop_id=req.photoshop.id or 0,
        image_url=image.image_url,
        image_type=image.image_type,
        started_at=started_at,
        finished_at=finished_at,
        logs="",
    )
