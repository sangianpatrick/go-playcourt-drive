.PHONY: test cover


test:
		go test -v -coverprofile=coverage.out -covermode=atomic ./...

cover: test
	go tool cover -func=coverage.out &&\
		go tool cover -html=coverage.out -o coverage.html
