make all: sqlcgen build

sqlcgen:
	cd db && sqlc generate

build: */*.go go.sum go.mod
	go build -o bin/cafe_manager main.go

test:
	go test ./...

cover-test:
	go test -coverprofile=coverage.out ./...
	gocover-cobertura < coverage.out > coverage.xml
