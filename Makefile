include local.env

get-deps:
	go mod init github.com/Zasedatelev/Todo.git
	go mod tidy


migrateUp:
	migrate -path migrations -database "postgres://$(PG_USER):$(PG_PASSWORD)@localhost:$(PG_PORT)/$(PG_NAME)?sslmode=disable" up

migrateDown:
	migrate -path migrations -database "postgres://$(PG_USER):$(PG_PASSWORD)@localhost:$(PG_PORT)/$(PG_NAME)?sslmode=disable" down

start:
	go run cmd/todo/main.go  