BINARY_NAME=weflow

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows.exe main.go
run:
	./${BINARY_NAME}-darwin

build_and_run:build run

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows.exe
