DB_URL=postgresql://root:root@localhost:5432/dbdata?sslmode=disable

develop:
	echo "Starting docker environment"
	docker-compose -f docker-compose.yml up --build

createdb:
	docker exec -it postgres createdb --username=root --owner=root $(DB_NAME)

dropdb:
	docker exec -it postgres dropdb $(DB_NAME)

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go
	
.PHONY: develop