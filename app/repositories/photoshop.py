from sqlalchemy.orm import Session

from app.entities.photoshop import PhotoshopElement, PhotoshopFile


class PhotoshopFileRepository:
    @staticmethod
    def find_all(db: Session) -> list[PhotoshopFile]:
        return db.query(PhotoshopFile).all()

    @staticmethod
    def find_by_id(db: Session, id: int) -> PhotoshopFile | None:
        return db.query(PhotoshopFile).get(id)

    @staticmethod
    def save(db: Session, photoshop_file: PhotoshopFile) -> PhotoshopFile:
        if photoshop_file.id is not None:
            db.merge(photoshop_file)
        else:
            db.add(photoshop_file)
        db.commit()
        return photoshop_file


class PhotoshopElementRepository:
    @staticmethod
    def find_all(db: Session) -> list[PhotoshopElement]:
        return db.query(PhotoshopElement).all()

    @staticmethod
    def find_by_id(db: Session, id: int) -> PhotoshopElement | None:
        return db.query(PhotoshopElement).get(id)

    @staticmethod
    def save(db: Session, photoshop_element: PhotoshopElement) -> PhotoshopElement:
        if photoshop_element.id is not None:
            db.merge(photoshop_element)
        else:
            db.add(photoshop_element)
        db.commit()
        return photoshop_element
