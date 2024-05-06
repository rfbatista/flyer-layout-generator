
from typing import Optional
from pydantic import BaseModel

from app.entities.componente import Componente


class DesignTemplateRegion(BaseModel):
    id: str
    xi: int
    xii: int
    yi: int
    yii: int
    component: Optional[Componente] = None

    def set_component(self, c: Componente):
        self.component = c

    def bbox(self):
        return ((self.xi, self.yi), (self.xii, self.yii))

    def size(self):
        return (self.xii - self.xi, self.yii - self.yi)

    def width(self):
        return self.size()[0]

    def height(self):
        return self.size()[1]

