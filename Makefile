BINARY_NAME=web-server

build:
	go build -o ${BINARY_NAME} cmd/main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}-darwin 2> /dev/null
	rm ${BINARY_NAME}-linux 2> /dev/null
	rm ${BINARY_NAME}-windows 2> /dev/null