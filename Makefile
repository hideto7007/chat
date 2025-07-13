up:
	docker-compose -f docker-compose.yml -p chat up -d
	(cd apps/frontend && ls && npm watch:swagger)
	wait


up-nocache:
	docker-compose -p chat down -v
	docker volume rm $(docker volume ls -q) || true
	docker-compose -f docker-compose.yml -p chat up --build -d
	(cd apps/frontend && npm watch:swagger && cd apps/backend/cmd/reset && go run main.go)
	wait

migrate:
	cd apps/backend/cmd/migrate && go run main.go --dbhost=localhost

swag-generate:
	cd apps/backend && swag init

reset:
	cd apps/backend/cmd/reset && go run main.go --dbhost=localhost

build:
	cd apps/backend && go build -o chat main.go