from datetime import datetime, timezone
from PIL import Image, ImageDraw


from app.entities.prancheta import DesignPrancheta
from app.upload_image import upload_image
from psd_tools import PSDImage
from pydantic import BaseModel


class BuildImageRequest(BaseModel):
    design_file: str
    prancheta: DesignPrancheta


class BuildImageResponse(BaseModel):
    image_url: str


def build_image(req: BuildImageRequest):
    print(req.prancheta.grid)
    psd = PSDImage.open(req.design_file)
    print(req.prancheta)
    img = Image.new("RGB", (req.prancheta.width, req.prancheta.height), "black")
    if req.prancheta.background is not None:
        req.prancheta.background.index_elements(psd)
        req.prancheta.background.draw_in_image(img, log=True)
    for c in req.prancheta.components:
        print(c)
        c.index_elements(psd)
        c.draw_in_image(img, log=True)

    created_at = datetime.now(timezone.utc)
    draw = ImageDraw.Draw(img)
    if req.prancheta.grid is not None:
        if req.prancheta.grid.regions is not None and len(req.prancheta.grid.rectangle) > 0:
            for i in req.prancheta.grid.regions:
                draw.rectangle([(i.xi,i.yi),(i.xii,i.yii)], outline ="red")

    image_url = upload_image(img, str(created_at))
    return BuildImageResponse(
        image_url=image_url,
    )
