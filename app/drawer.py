from datetime import datetime, timezone
from typing import List

from PIL import Image, ImageDraw
from psd_tools import PSDImage
from pydantic import BaseModel
from app.entities.componente import Componente
from app.requests import GenerateDesignRequest
from app.upload_image import upload_image


class Design(BaseModel):
    image_url: str
    image_type: str


class Renderer:
    def render_png(
        self, components: List[Componente], req: GenerateDesignRequest 
    ) -> Design:
        psd = PSDImage.open(req.photoshop.filepath)
        for c in components:
            c.index_elements(psd)
        img = Image.new("RGB", (req.template.width, req.template.height), "black")
        draw = ImageDraw.Draw(img)
        # draw lines for testing
        # for i in range(4):
        #     draw.line([(150*i,0), (150*i, 200)], fill=None, width=3)
        for c in components:
            c.draw_in_image(img)
        created_at = datetime.now(timezone.utc)
        image_url = upload_image(img, str(created_at))
        return Design(
            image_url=image_url,
            image_type="png",
        )


renderer = Renderer()
