from copy import deepcopy
from typing import List
from pydantic import BaseModel

from app.entities.componente import Componente

class ComponentHistory(BaseModel):
    components: List[Componente] = []

    def add(self, c: Componente):
        self.components.append(deepcopy(c))

