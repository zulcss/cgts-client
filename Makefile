BINDIR := bin
CLI := system

.PHONY: build
build:
	go build -o $(BINDIR)/$(CLI)

.PHONY: test
test: build unit-tests lint

.PHONY: unit-tests
unit-tests:
	go test

.PHONY: clean
clean:
	rm -rf bin
	go clean

