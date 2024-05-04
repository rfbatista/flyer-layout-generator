from sqlalchemy import select
from sqlalchemy.orm.session import Session

from app.entities.photoshop import DesignElement


def list_photoshop_element(db: Session, ind: int):
    stmt = select(DesignElement).filter(DesignElement.photoshop_id == ind)
    data = db.scalars(stmt).all()
    return data
