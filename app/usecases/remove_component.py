from typing import List
from fastapi import Depends
from pydantic import BaseModel
from sqlalchemy.orm import Session

from app.entities.photoshop import DesignElement


class RemoveComponentsUseCaseRequest(BaseModel):
    elements_id: List[int]


def remove_components_usecase(req: RemoveComponentsUseCaseRequest, db: Session):
    db.query(DesignElement).filter(DesignElement.id.in_(req.elements_id)).update(
        {
            "component_id": None,
            "is_background": False,
        }
    )
    db.commit()
    db.flush()
    return
