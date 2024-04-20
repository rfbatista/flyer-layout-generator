from sqlalchemy import select
from sqlalchemy.orm import Session

from app.entities.gen_request import GenerationRequestImage


def list_generated_images_usecase(id: int, db: Session):
    stmt = select(GenerationRequestImage).filter(
        GenerationRequestImage.photoshop_id == id
    )
    data = db.scalars(stmt).all()
    return data


def list_generated_images_by_request_usecase(id: int, db: Session):
    stmt = select(GenerationRequestImage).filter(
        GenerationRequestImage.request_id == id
    )
    data = db.scalars(stmt).all()
    return data
