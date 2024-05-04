from sqlalchemy import select
from sqlalchemy.orm import join
from sqlalchemy.orm.session import Session

from app.entities.photoshop import DesignElement, PhotoshopFile


def list_photoshop_files(db: Session, limit=10, skip=0):
    stmt = select(PhotoshopFile).offset(skip).limit(limit)
    data = db.scalars(stmt).all()
    return data
