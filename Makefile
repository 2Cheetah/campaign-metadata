compose-build:
	docker compose build

compose-dev-up:
	docker compose -f compose.yaml -f compose-dev.yaml up

compose-dev-down:
	docker compose -f compose.yaml -f compose-dev.yaml down
