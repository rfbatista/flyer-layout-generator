from datetime import datetime, timezone
from typing import List
from PIL import Image, ImageDraw


from app.entities.prancheta import DesignPrancheta
from app.upload_image import upload_image
from psd_tools import PSDImage
from pydantic import BaseModel


class BuildImageRequest(BaseModel):
    design_file: str
    prancheta: DesignPrancheta


class ImageResponseElements(BaseModel):
    element_id: int
    image_url: str


class BuildImageResponse(BaseModel):
    image_url: str
    elements: List[ImageResponseElements]


def build_image(req: BuildImageRequest):
    print(req.prancheta.grid)
    psd = PSDImage.open(req.design_file)
    print(req.prancheta)
    img = Image.new("RGB", (req.prancheta.width, req.prancheta.height), "black")
    if req.prancheta.background is not None:
        req.prancheta.background.index_elements(psd)
        req.prancheta.background.draw_in_image(img, log=True)
    elements = []
    for c in req.prancheta.components:
        c.index_elements(psd)
        for elem in c._items:
            im = elem.image()
            size = (100, 100)
            for celem in c.elements:
                if celem.id == elem.id:
                    size = elem.size()
            print("\n")
            print(size)
            print("\n")
            im = im.resize((size[0],size[1]), Image.Resampling.LANCZOS)
            image_url = upload_image(im, "id{}".format(elem.id))
            elements.append(
                ImageResponseElements(
                    element_id=elem.id,
                    image_url=image_url,
                )
            )
        c.draw_in_image(img, log=True)

    created_at = datetime.now(timezone.utc)
    draw = ImageDraw.Draw(img)
    if req.prancheta.grid is not None:
        if (
            req.prancheta.grid.regions is not None
            and len(req.prancheta.grid.regions) > 0
        ):
            for i in req.prancheta.grid.regions:
                draw.rectangle([(i.xi, i.yi), (i.xii, i.yii)], outline="red")

    image_url = upload_image(img, str(created_at))
    return BuildImageResponse(
        image_url=image_url,
        elements=elements,
    )
