include ./env
export $(shell sed 's/=.*//' ./env)

test:
	go test ./...

run-appserver:
	go run ./app/web/main.go



.PHONY: all
