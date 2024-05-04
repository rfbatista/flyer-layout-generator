from typing import List

from pydantic import BaseModel

from app.entities.photoshop import DesignElement


class TemplateDistortion(BaseModel):
    x: int
    y: int


class TemplateSlotsPosition(BaseModel):
    xi: int
    yi: int
    width: int
    height: int


class Template(BaseModel):
    id: int
    type: str
    distortion: TemplateDistortion
    slots_positions: List[TemplateSlotsPosition]


class Photoshop(BaseModel):
    id: int
    name: str
    filepath: str
    image_path: str
    width: int
    height: int
    created_at: str


class PhotoshopComponent(BaseModel):
    id: str
    elements: List[DesignElement]


class GenerationRequest(BaseModel):
    template: Template
    photoshop: Photoshop
    components: List[PhotoshopComponent]
    elements: List[DesignElement]


class GenerationResult(BaseModel):
    request_id: str
    photoshop_id: int
    template_id: int
    image_url: str


def DistortionGenerationProcessor(req: GenerationRequest):
    return
