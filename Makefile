BINARY=gcase
BIN_DIR=bin

build:
	mkdir -p ${BIN_DIR}
	go build -o ${BIN_DIR}/${BINARY} .

run: build
	@./${BIN_DIR}/${BINARY}

fmt:
	go fmt ./...

test:
	go test ./...

test-verbose:
	go test -v ./...

clean:
	rm -f ./${BIN_DIR}/${BINARY}

tidy:
	go mod tidy
