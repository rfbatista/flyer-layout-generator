from typing import List
from pydantic import BaseModel

from app.dto.photoshop_element_dto import PhotoshopElementDto


class ComponentDto(BaseModel):
    id: int
    elements: List[PhotoshopElementDto]
