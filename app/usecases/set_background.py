from typing import List
from pydantic import BaseModel
from sqlalchemy.orm import Session

from app.entities.photoshop import PhotoshopElement


class SetBackgroundRequest(BaseModel):
    elements_id: List[int]

def set_background_usecase(req: SetBackgroundRequest, db: Session):
    db.query(PhotoshopElement).filter(PhotoshopElement.id.in_(req.elements_id)).update(
        {
            "component_color": None,
            "component_id": None,
            "is_background": True,
        }
    )
    db.commit()
    db.flush()
    return
