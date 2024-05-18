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


def expand_and_center(base_image: Image, base_width, base_height, target_width, target_height, target_image: Image):
    
    # Calculate the aspect ratio of the target image
    target_aspect_ratio = target_width / target_height
    
    # Calculate the new dimensions of the target image to fill the base image
    if (base_width / base_height) > target_aspect_ratio:
        new_width = base_width
        new_height = int(base_width / target_aspect_ratio)
    else:
        new_width = int(base_height * target_aspect_ratio)
        new_height = base_height
    
    # Resize the target image to the new dimensions
    resized_target_image = target_image.resize((new_width, new_height), Image.LACZOS)
    
    # Create a new image with the size of the base image and a white background
    new_image = Image.new('RGB', (base_width, base_height), (255, 255, 255))
    
    # Calculate the position to paste the resized target image to center it
    x_offset = (base_width - new_width) // 2
    y_offset = (base_height - new_height) // 2
    
    # Paste the resized target image onto the new image
    new_image.paste(resized_target_image, (x_offset, y_offset))
    
    return new_image

class Renderer:
    def render_png(
        self, components: List[Componente], req: GenerateDesignRequest, background: Componente
    ) -> Design:
        psd = PSDImage.open(req.photoshop.filepath)
        for c in components:
            c.index_elements(psd)
        background.index_elements(psd)
        img = Image.new("RGB", (req.template.width, req.template.height), "black")
        background.draw_in_image(img)
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
