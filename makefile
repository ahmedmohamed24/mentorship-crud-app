run-local:
	docker-compose -f docker-compose.yaml up -d

run-live:
	docker-compose -f docker-compose-live.yaml up -d

run-server:
	go run main.go serve

migrate:
	go run main.go migrate

seed:
	go run main.go seed

down-local:
	docker-compose -f docker-compose.yaml down

down-live:
	docker-compose -f docker-compose-live.yaml down