import io
import requests
from PIL.Image import Image


endpoint_url = "http://localhost:8000/api/v1/images"
# endpoint_url = "http://app-dev:8000/api/v1/images"


def upload_image(img: Image, name: str)->str:
    image_data = None
    with io.BytesIO() as output:
        img.save(output, format="PNG", quality=95)
        image_data = output.getvalue()
    # Endpoint URL where you want to send the image

    # Create a dictionary containing the file data
    files = {"file": ("image.png", image_data, "image/png"), "filename": name}

    # Send the POST request with the image file
    response = requests.post(endpoint_url, files=files, data={"filename": name})

    # Check the response status
    if response.status_code == 200:
        print("Image uploaded successfully.")
        data = response.json()
        return data["image_url"]
    else:
        print("Failed to upload image:", response.text)
