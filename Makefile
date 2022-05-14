migrations:
	goose -dir .data/migrations postgres "postgres://root:root@localhost:5432/awesomeProject?sslmode=disable" up
up:
	docker-compose up -d
down:
	docker-compose down
env:
	cp ./.env.example ./.env
http:
	go run main.go httpserver