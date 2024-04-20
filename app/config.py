import os

from dotenv import dotenv_values, load_dotenv
from pydantic import BaseModel

load_dotenv()


class Config:
    def __init__(self, conf):
        self._config = conf

    @property
    def photoshop_files_path(self):
        return self._config["photoshop_files_path"]

    @property
    def dist_path(self):
        return self._config["dist_path"]

    @property
    def log_filename(self):
        return self._config["log_file"]

app_config = Config(
    {
        "dist_path": os.getenv("DIST_PATH"),
        "photoshop_files_path": os.getenv("PHOTOSHOP_FILES_PATH"),
        "log_file": os.getenv("LOG_FILE_PATH"),
    }
)
