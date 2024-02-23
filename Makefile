build:
	docker compose -f docker-compose.yml build --force-rm
run:
	docker compose -f docker-compose.yml up

