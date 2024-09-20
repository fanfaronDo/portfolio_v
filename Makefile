run:
	@go run ./cmd/app/main.go

build:
	go build -o portfolio_run ./cmd/app/main.go

build_test:
	go build -o test ./cmd/test/main.go