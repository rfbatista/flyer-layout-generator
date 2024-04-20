from collections.abc import Iterable
from contextlib import asynccontextmanager

from fastapi import BackgroundTasks, Depends, FastAPI, File, HTTPException, UploadFile
from fastapi.middleware.cors import CORSMiddleware
from fastapi.staticfiles import StaticFiles
from psd_tools import PSDImage
from sqlalchemy.orm.session import Session

from app.db import get_db
from app.entities.image import Image
from app.entities.photoshop import Elemento
from app.logger import logger
from app.repositories.photoshop import PhotoshopFileRepository
from app.usecases.create_component import (
    CreateComponentUseCaseRequest,
    create_component_usecase,
)
from app.usecases.create_template import CreateTemplateRequest, create_template
from app.usecases.generate_designs import (
    GenerateDesignRequest,
    design_generation_usecase,
)
from app.usecases.list_generated_images import list_generated_images_by_request_usecase, list_generated_images_usecase
from app.usecases.list_photoshop_elements import list_photoshop_element
from app.usecases.list_photoshop_file import list_photoshop_files
from app.usecases.list_templates import list_templates
from app.usecases.remove_component import (
    RemoveComponentsUseCaseRequest,
    remove_components_usecase,
)
from app.usecases.save_photoshop_file import save_photoshop_file
from app.usecases.set_background import SetBackgroundRequest, set_background_usecase


@asynccontextmanager
async def startup_event(_: FastAPI):
    logger.info("starting up application")
    yield


app = FastAPI(lifespan=startup_event)
app.mount("/static", StaticFiles(directory="dist"), name="static")

origins = [
    "http://localhost:5173",
]


app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.post("/v1/design")
def generate_layouts():
    return


@app.get("/v1/generation/status")
def generation_status():
    return


@app.get("/api/v1/photoshop/{id}/design")
def list_generated_images(id:int, db: Session = Depends(get_db)):
    try:
        return list_generated_images_usecase(id, db)
    except Exception as e:
        logger.exception("failed to list images")
        raise HTTPException(status_code=500, detail="internal server error \n %s" % (e))

@app.get("/api/v1/request/{id}/design")
def list_generated_images_by_request_api(id:int, db: Session = Depends(get_db)):
    try:
        return list_generated_images_by_request_usecase(id, db)
    except Exception as e:
        logger.exception("failed to list images by request")
        raise HTTPException(status_code=500, detail="internal server error \n %s" % (e))

@app.post("/api/v1/template")
def create_template_endpoint(req: CreateTemplateRequest, db: Session = Depends(get_db)):
    return create_template(db, req)


@app.get("/api/v1/template")
def list_templates_endpoint(
    skip: int = 0, limit: int = 10, db: Session = Depends(get_db)
):
    return list_templates(db, skip, limit)


@app.post("/api/v1/component")
def create_component_api(
    req: CreateComponentUseCaseRequest, db: Session = Depends(get_db)
):
    try:
        return create_component_usecase(req, db)
    except Exception as e:
        logger.exception("failed to create component")
        raise HTTPException(status_code=500, detail="internal server error \n %s" % (e))


@app.post("/api/v1/component/remove")
def remove_components_api(
    req: RemoveComponentsUseCaseRequest, db: Session = Depends(get_db)
):
    try:
        return remove_components_usecase(req, db)
    except Exception as e:
        logger.exception("failed to remove components")
        raise HTTPException(status_code=500, detail="internal server error \n %s" % (e))


@app.post("/api/v1/background")
def set_background_api(req: SetBackgroundRequest, db: Session = Depends(get_db)):
    try:
        return set_background_usecase(req, db)
    except Exception as e:
        logger.exception("failed to set background")
        raise HTTPException(status_code=500, detail="internal server error \n %s" % (e))


@app.post("/api/v1/design")
def design_generation_api(
    req: GenerateDesignRequest,
    background_tasks: BackgroundTasks,
    db: Session = Depends(get_db),
):
    try:
        return design_generation_usecase(req, db, background_tasks)
    except Exception as e:
        logger.exception("failed to generate design")
        raise HTTPException(status_code=500, detail="internal server error \n %s" % (e))


@app.get("/api/v1/photoshop")
def get_images_from_photoshop(
    skip: int = 0, limit: int = 10, db: Session = Depends(get_db)
):
    return list_photoshop_files(db, limit, skip)


@app.get("/api/v1/photoshop/{id}/elements")
def get_photoshop_elements(id: int, db: Session = Depends(get_db)):
    try:
        return list_photoshop_element(db, id)
    except Exception as e:
        logger.exception("failed to list photoshop elements")
        raise HTTPException(status_code=500, detail="internal server error \n %s" % (e))


@app.get("/api/v1/photoshop/{photoshop_id}")
def get_file(photoshop_id: int, db: Session = Depends(get_db)):
    return


@app.get("/api/v1/photoshop/{photoshop_id}/images")
def get_file_images(photoshop_id: int, db: Session = Depends(get_db)):
    data = PhotoshopFileRepository.find_by_id(db, photoshop_id)
    if data:
        psd = PSDImage.open(data.filepath)
        psd.size
        return [Image(psd)]
    else:
        return data


@app.post("/api/v1/photoshop")
def save_file(file: UploadFile, db: Session = Depends(get_db)):
    return save_photoshop_file(file, db)
