from sqlalchemy import Column, Integer
from app.db import Base


class Position(Base):
    id: int = Column(Integer, primary_key=True, index=True)
    xi: int = Column(Integer, primary_key=True, index=True)
    yi: int = Column(Integer, primary_key=True, index=True)
    xii: int = Column(Integer, primary_key=True, index=True)
    yii: int = Column(Integer, primary_key=True, index=True)


class Template(Base):
    id: int = Column(Integer, primary_key=True, index=True)
