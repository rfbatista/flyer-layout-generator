run:
	PYTHONPATH=. uvicorn app.main:app --reload
migrate:
	PYTHONPATH=. alembic revision --autogenerate -m $(msg)
upgrade:
	alembic upgrade head
