from sqlalchemy import select
from sqlalchemy.orm import Session

from app.entities.template import Position, Template


def list_templates(db: Session, skip: int, limit: int):
    stmt = select(Template)
    result = db.scalars(stmt).all()
    dtos = []
    for item in result:
        pos_stmt = select(Position).filter(Position.template_id == item.id)
        pos_result = db.scalars(pos_stmt).all()
        dtos.append(Template.from_db(item, pos_result))
    return dtos
