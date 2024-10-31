run:
	@go run ./cmd/app/main.go

build:
	go build -o portfolio_run ./cmd/app/main.go

build_test:
	go build -o test ./cmd/test/main.go

ttt:
	export PATH=$(PATH):/usr/local/go/bin
	echo $(go env)

# install_migrate:
# 	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# 	@echo "alais migrate=$(go env GOPATH)/bin/migrate"

migrationup: install_migrate
	@migrate -source file://migrations -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST))/$(MYSQL_DATABASE)" up
 
# migrationdown: export_tools
# 	migrate -source file://migrations -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(192.168.56.2:3306)/$(MYSQL_DATABASE)" down
