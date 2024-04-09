import uuid

from fastapi import Depends, FastAPI, File, UploadFile
from fastapi.middleware.cors import CORSMiddleware
from fastapi.staticfiles import StaticFiles
from psd_tools import PSDImage
from sqlalchemy.orm.session import Session

from app.config import app_config
from app.db import get_db
from app.entities.image import Image
from app.entities.photoshop import PhotoshopFile
from app.repositories.photoshop import PhotoshopFileRepository

app = FastAPI()
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


@app.get("/v1/generation")
def get_generation_result():
    return


@app.get("/api/v1/photoshop")
def get_images_from_photoshop(db: Session = Depends(get_db)):
    data = PhotoshopFileRepository.find_all(db)
    print(data)
    return data

@app.get("/api/v1/photoshop/{photoshop_id}")
def get_file(photoshop_id: int, db: Session = Depends(get_db)):
    data = PhotoshopFileRepository.find_by_id(db, photoshop_id)
    if not data: return
    psd = PSDImage.open(data.filepath)
    if not psd: return
    items = []
    def index_elements(layer, pre = ""):
        if not isinstance(layer, Iterable):
          return
        for idx, layer in enumerate(layer):
          if layer.layer_id in self.items:
            items.append(Elemento(layer))
        index_elements(layer, pre + '\t')
    index_elements(psd)
    return items
        return [Image(psd)]

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
    try:
        unique_id = uuid.uuid4()
        path = "%s/%s" % (app_config["PHOTOSHOP_FILES_PATH"], unique_id)
        with open(path, "wb") as f:
            f.write(file.file.read())
        psd = PSDImage.open(path)
        photoshopfile = PhotoshopFile(
            filename=file.filename,
            filepath=path,
            width=psd.width,
            height=psd.height,
        )
        PhotoshopFileRepository.save(db, photoshopfile)
        return {
            "file": file.filename,
            "content": file.content_type,
            "path": path,
        }
    except Exception as e:
        return {"message": e.args}
