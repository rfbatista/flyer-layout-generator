import os
import uuid

from psd_tools import PSDImage

dist_path = os.getenv("DIST_PATH")


class Image:
    def __init__(self, psd: PSDImage) -> None:
        self.bbox = psd.bbox
        filename = uuid.uuid4()
        filepath = "%s/%s.jpg" % (dist_path, filename)
        img = psd.composite()
        if img:
            img.save(filepath)
        self.image = "/static/%s.jpg" % (filename)
