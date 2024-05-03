from contextlib import asynccontextmanager

from pydantic import BaseModel
from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from fastapi.staticfiles import StaticFiles

from app.generate_designs import GenerateDesignRequest, generate_design
from app.process_photoshop_file import process_photoshop_file
from .logger import logger


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

class ProcessFileRequest(BaseModel):
    filepath: str

@app.post("/api/v1/photoshop")
def save_file(req: ProcessFileRequest):
    try:
        return process_photoshop_file(req.filepath)
    except Exception as e:
        logger.exception("failed to set background")
        raise HTTPException(status_code=500, detail="internal server error \n %s" % (e))


@app.post("/api/v1/generate/distortion")
def generate_design_api(req: GenerateDesignRequest):
    try:
        return generate_design(req)
    except Exception as e:
        logger.exception("failed to generate design")
        raise HTTPException(status_code=500, detail="internal server error \n %s" % (e))
