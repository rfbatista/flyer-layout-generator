from datetime import datetime
from typing import List

from pydantic import BaseModel
from app.entities.componente import Componente
from app.entities.photoshop import PhotoshopElement, PhotoshopFile
from app.entities.template import DesignTemplate


class GenerateDesignRequest(BaseModel):
    photoshop: PhotoshopFile
    template: DesignTemplate
    elements: List[PhotoshopElement]
    components: List[Componente]


class GenerateDesignResult(BaseModel):
    photoshop_id: int
    image_url: str
    image_type: str
    started_at: datetime 
    finished_at: datetime
    logs: str
