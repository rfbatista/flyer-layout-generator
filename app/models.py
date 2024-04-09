import peewee
from app.entities.photoshop import PhotoshopFile


if __name__ == "__main__":
    try:
        PhotoshopFile.create_table()
        print("Tabela 'PhotoshopFile' criada com sucesso!")
    except peewee.OperationalError:
        print("Tabela 'PhotoshopFile' ja existe!")
