include ./env
export $(shell sed 's/=.*//' ./env)

test:
	go test ./...

run-appserver:
	go run ./app/web/main.go

cli-create-table:
	go run ./app/cli/main.go create-table

.PHONY: all
