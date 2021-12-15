demo.up:
	cd internal/example/ && go run .

compose.up:
	docker-compose up -d
compose.down:
	docker-compose down

