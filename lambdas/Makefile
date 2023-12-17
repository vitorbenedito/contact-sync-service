.PHONY: build clean deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/health health/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/contacts-sync contacts-sync/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/contacts-sync-parallel contacts-sync-parallel/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
