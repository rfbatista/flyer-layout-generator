from typing import Iterable, List
from fastapi import BackgroundTasks, HTTPException
from psd_tools import PSDImage
from pydantic import BaseModel
from sqlalchemy import select
from sqlalchemy.orm import Session
import uuid

from app.entities.componente import Componente
from app.config import app_config
from app.entities.gen_request import GenerationRequest, GenerationRequestImage
from app.entities.photoshop import DesignElement, PhotoshopFile
from app.entities.prancheta import Prancheta
from app.entities.template import Template
from app.logger import logger


class BackgroundPosition(BaseModel):
    xi: int
    yi: int


class TemplateDetails(BaseModel):
    id: int
    background_position: BackgroundPosition


class GenerateDesignRequest(BaseModel):
    templates: List[TemplateDetails]
    photoshop_id: int


def design_generation_usecase(
    req: GenerateDesignRequest, db: Session, background_tasks: BackgroundTasks
):
    stmt = select(PhotoshopFile).filter(PhotoshopFile.id == req.photoshop_id)
    photoshop_file = db.scalars(stmt).first()
    if photoshop_file is None:
        raise HTTPException(400, detail="photoshop file was not found in database")
    psd = PSDImage.open(photoshop_file.filepath)
    request = GenerationRequest(status="started", photoshop_id=req.photoshop_id)
    db.add(request)
    db.commit()
    db.flush()
    db.expunge(photoshop_file)
    background_tasks.add_task(
        create_design, req.photoshop_id, psd, req.templates[0], request.id, db
    )
    return request


def create_design(
    photoshop_id: int,
    psd: PSDImage,
    template_details: TemplateDetails,
    gen_request_id: int,
    db: Session,
):
    logger.info("starting to create design")
    stmt = (
        select(Template)
        .join(Template.positions)
        .filter(Template.id == template_details.id)
    )
    template = db.scalars(stmt).first()
    if template is None:
        logger.error("failed to find template %s" % (template_details.id))
        return
    logger.info("fetching photoshop elements")
    stmt = select(DesignElement).filter(
        DesignElement.photoshop_id == photoshop_id
    )
    elements = db.scalars(stmt).all()
    elements_group = {}
    background_elem = []
    for elem in elements:
        if elem.is_component():
            if elem.component_id in elements_group.keys():
                elements_group[elem.component_id].append(elem.layer_id)
            else:
                elements_group[elem.component_id] = [elem.layer_id]
        if elem.is_background:
            background_elem.append(elem.layer_id)
    components = []
    for idx in elements_group:
        components.append(Componente(elements_group[idx]))
    background = Componente()
    for idx in background_elem:
        background.add_element(idx)
    background.index_elements(psd)
    prancheta = Prancheta(template.width, template.height)
    prancheta.set_background(background)
    for component in components:
        component.index_elements(psd)
        prancheta.add_componente(component)
    logger.info("creating image")
    img = prancheta.image_from_template(template)
    filename = "%s.png" % (uuid.uuid4())
    filepath = "%s%s" % (app_config.dist_path, filename)
    img.save(filepath)
    image_req = GenerationRequestImage(
        filename=filename,
        filepath=filepath,
        photoshop_id=photoshop_id,
        template_id=template.id,
        request_id=gen_request_id,
    )
    db.add(image_req)
    db.commit()
    db.flush()
