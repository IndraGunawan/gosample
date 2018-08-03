include ./env
export $(shell sed 's/=.*//' ./env)

run-appserver:
	go run ./app/web/main.go

test:
	go test ./...

.PHONY: test
