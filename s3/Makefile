.PHONY: build clean deploy

build:
	dep ensure -v
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello main.go
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bin/main main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/main main.go
	GOARCH=arm64 GOOS=linux go build -tags lambda.norpc -o ./bin/bootstrap
	while true; do [[ -f ../bin/bootstrap ]] || GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o ../bin/bootstrap ; sleep 0.1 ; done
clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
