run:
	@go run ./cmd/app/main.go

build:
	go build -o portfolio_run ./cmd/app/main.go

build_test:
	go build -o test ./cmd/test/main.go

migrationup:
	migrate -source file://migrations -database "mysql://mysql:root@tcp(192.168.56.103:3306)/portfolio" up

migrationdown:
	migrate -source file://migrations -database "mysql://mysql:root@tcp(192.168.56.103:3306)/portfolio" down
