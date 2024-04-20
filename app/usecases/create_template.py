from typing import List
from fastapi import HTTPException
from pydantic import BaseModel
from sqlalchemy.orm import Session

from app.logger import logger
from app.entities.template import Position, Template


class CreateTemplateRequestPoint(BaseModel):
    xi: int
    yi: int
    width: int
    height: int


class CreateTemplateRequest(BaseModel):
    name: str
    width: int
    height: int
    positions: List[CreateTemplateRequestPoint]


def create_template(db: Session, req: CreateTemplateRequest):
    try:
        template = Template(name=req.name,width=req.width,height=req.height)
        db.add(template)
        db.commit()
        db.flush()
        points = []
        for position in req.positions:
            points.append(
                Position(
                    xi=position.xi,
                    yi=position.yi,
                    width=position.width,
                    height=position.height,
                    template_id=template.id,
                )
            )
        db.add_all(points)
        db.commit()
        db.flush()
        return template
    except Exception as e:
        logger.exception("failed to save photoshop file")
        raise HTTPException(status_code=500, detail="internal server error \n %s" % (e))
