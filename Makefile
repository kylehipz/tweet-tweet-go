.PHONY: dev

dev/build:
	docker compose -f deployments/compose/compose.yaml up --build -d

dev/up:
	docker compose -f deployments/compose/compose.yaml up -d

dev/down:
	docker compose -f deployments/compose/compose.yaml down
