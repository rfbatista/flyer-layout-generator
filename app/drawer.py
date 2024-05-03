from datetime import datetime, timezone
from typing import List
import uuid

from PIL import Image
from psd_tools import PSDImage
from pydantic import BaseModel
from app.entities.componente import Componente
from app.config import app_config
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
        for c in components:
            c.draw_in_image(img)
        created_at = datetime.now(timezone.utc)
        image_url = upload_image(img, str(created_at))
        return Design(
            image_url=image_url,
            image_type="png",
        )


renderer = Renderer()
