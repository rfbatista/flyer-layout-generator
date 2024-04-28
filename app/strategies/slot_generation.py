from typing import List
from pydantic import BaseModel

from app.entities.componente import Componente


class SlotGenerationRequest(BaseModel):
    photoshop_url: str
    components: List[Componente]
