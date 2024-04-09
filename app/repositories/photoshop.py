from sqlalchemy.orm import Session

from app.entities.photoshop import PhotoshopFile


class PhotoshopFileRepository:
    @staticmethod
    def find_all(db: Session) -> list[PhotoshopFile]:
        return db.query(PhotoshopFile).all()

    @staticmethod
    def find_by_id(db: Session, id: int) -> PhotoshopFile | None:
        return db.query(PhotoshopFile).get(id)

    @staticmethod
    def save(db: Session, photoshop_file: PhotoshopFile) -> PhotoshopFile:
        if photoshop_file.id:
            db.merge(photoshop_file)
        else:
            db.add(photoshop_file)
        db.commit()
        return photoshop_file
