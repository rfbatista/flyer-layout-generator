from typing import List
from pydantic import BaseModel
from sqlalchemy.orm import Session

from app.entities.photoshop import DesignElement


class CreateComponentUseCaseRequest(BaseModel):
    elements_id: List[int]
    color: str
    component_id: str


def create_component_usecase(req: CreateComponentUseCaseRequest, db: Session):
    db.query(DesignElement).filter(DesignElement.id.in_(req.elements_id)).update(
        {
            "component_color": req.color,
            "component_id": req.component_id,
            "is_background": False,
        }
    )
    db.commit()
    db.flush()
    return
