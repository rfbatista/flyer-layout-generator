import string
import random

alphabet = string.ascii_lowercase + string.digits

def create_id()->str:
    return "".join(random.choices(alphabet, k=8))
