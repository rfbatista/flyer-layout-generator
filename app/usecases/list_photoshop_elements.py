from sqlalchemy import select
from sqlalchemy.orm.session import Session

from app.entities.photoshop import PhotoshopElement


def list_photoshop_element(db: Session, ind: int):
    stmt = select(PhotoshopElement).filter(PhotoshopElement.photoshop_id == ind)
    data = db.scalars(stmt).all()
    return data
