from sqlalchemy import Column, DateTime, ForeignKey, Integer, String, func
from sqlalchemy.orm import Mapped, mapped_column

from app.db import Base

class GenerationRequest(Base):
    __tablename__ = "generation_request"

    id = Column(Integer, primary_key=True, index=True)
    status = mapped_column(String(10))
    log = mapped_column(String(50))
    photoshop_id = mapped_column(ForeignKey("photoshop_files.id"))


class GenerationRequestImage(Base):
    __tablename__ = "generation_request_image"

    id = Column(Integer, primary_key=True, index=True)
    filename = mapped_column(String(50))
    filepath = mapped_column(String(100))
    request_id = mapped_column(ForeignKey("generation_request.id"))
    photoshop_id= mapped_column(ForeignKey("photoshop_files.id"))
    template_id = mapped_column(ForeignKey("templates.id"))
    created_at = Column(DateTime(timezone=True), server_default=func.now())
