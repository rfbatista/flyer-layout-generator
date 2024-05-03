from datetime import datetime, timezone


from app.distortion import Dimension, distort_image_to
from app.drawer import renderer
from app.entities.prancheta import DesignPrancheta
from app.find_component_per_region import define_components_per_region
from app.position_components import position_components_in_regions
from app.requests import GenerateDesignRequest, GenerateDesignResult


def generate_design(req: GenerateDesignRequest):
    started_at = datetime.now(timezone.utc)
    prancheta = DesignPrancheta(
        template=req.template,
        components=req.components,
        width=req.photoshop.width,
        height=req.photoshop.height,
    )
    distorted_prancheta = distort_image_to(
        prancheta, Dimension(width=req.template.width, height=req.template.height)
    )
    regions = req.template.regions()
    regions_with_components = define_components_per_region(
        regions, distorted_prancheta.components
    )
    regions_with_components_positioned = position_components_in_regions(
        regions_with_components
    )
    componentes = [
        c.component
        for c in regions_with_components_positioned
        if c.component is not None
    ]
    image = renderer.render_png(componentes, req)
    finished_at = datetime.now(timezone.utc)
    return GenerateDesignResult(
        photoshop_id=req.photoshop.id or 0,
        image_url=image.image_url,
        image_type=image.image_type,
        started_at=started_at,
        finished_at=finished_at,
        logs="",
    )
