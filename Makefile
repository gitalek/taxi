lint:
	golangci-lint run

test:
	go test ./... -v -cover

test_calc:
	@cd ./calc/pkg && go test -v -cover
