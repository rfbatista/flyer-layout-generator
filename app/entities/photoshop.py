from sqlalchemy import Column, Integer, String

from app.db import Base

class PhotoshopFile(Base):
    __tablename__ = "photoshop_files"

    id: int = Column(Integer, primary_key=True, index=True)
    filename: str = Column(String(100), nullable=False)
    filepath: str = Column(String(255), nullable=False)
    width: int = Column(Integer)
    height: int = Column(Integer)
