lint:
	golangci-lint run

test:
	go test ./... -v -cover

test_calc:
	@cd ./calc/pkg && go test -v -cover

migrate:
	@cd ./internal/migrations && GOOSE_DRIVER=postgres GOOSE_DBSTRING="user=user password=password dbname=taxi_db sslmode=disable host=localhost port=54320" goose up

status:
	@cd ./internal/migrations && GOOSE_DRIVER=postgres GOOSE_DBSTRING="user=user password=password dbname=taxi_db sslmode=disable host=localhost port=54320" goose status
