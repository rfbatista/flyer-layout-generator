import logging
from app.config import app_config

logger = logging.getLogger(__name__)
logging.basicConfig(
    filename=app_config.log_filename,
    filemode="a",
    format="%(asctime)s %(name)-12s %(levelname)-8s %(message)s",
    datefmt="%m-%d %H:%M",
    level=logging.ERROR,
)
