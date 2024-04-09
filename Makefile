run:
	PYTHONPATH=. python app
migrate:
	PYTHONPATH=. alembic revision --autogenerate -m $(msg)
upgrade:
	alembic upgrade head
