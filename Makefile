install:
	go install .

deps:
	go get github.com/nats-io/nats
	go get github.com/fatih/color

dev-deps:
	go get -u github.com/golang/lint/golint
	go get -u github.com/smartystreets/goconvey/convey

test:
	go test -v ./...

lint:
	golint ./...
	go vet ./...
