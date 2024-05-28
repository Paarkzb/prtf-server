# go build -o ./dist/serverV2.exe ./cmd | .\dist\serverV2.exe

build:
	@echo building server
	go build -o ./dist/serverV2.exe ./cmd
	@echo server built

start: build
	@echo starting server
	.\dist\serverV2.exe
	@echo server started

migrateup:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable -verbose up $(v)

migratedown:
	migrate -path db/migrations -database postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable -verbose down $(v)

# migrate -path db/migrations -database postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable version