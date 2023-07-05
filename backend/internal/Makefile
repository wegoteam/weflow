BINARY_NAME=weflow
BINARY_FILE_PATH=bin

build:
	mkdir -p ${BINARY_FILE_PATH}
	GOARCH=amd64 GOOS=darwin go build -o ./${BINARY_FILE_PATH}/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ./${BINARY_FILE_PATH}/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ./${BINARY_FILE_PATH}/${BINARY_NAME}-windows.exe main.go
run:
	./${BINARY_NAME}-darwin

build_and_run:clean build run

clean:
	go clean
	rm -f -r ./${BINARY_FILE_PATH}
