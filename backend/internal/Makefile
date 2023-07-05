BINARY_NAME=weflow
BINARY_FILE_PATH=bin

build:
	mkdir -p ${BINARY_FILE_PATH}
	GOARCH=amd64 GOOS=linux go build -o ./${BINARY_FILE_PATH}/${BINARY_NAME}-linux main.go

build_windows:
	mkdir -p ${BINARY_FILE_PATH}
	GOARCH=amd64 GOOS=windows go build -o ./${BINARY_FILE_PATH}/${BINARY_NAME}-windows.exe main.go

build_mac:
	mkdir -p ${BINARY_FILE_PATH}
	GOARCH=amd64 GOOS=darwin go build -o ./${BINARY_FILE_PATH}/${BINARY_NAME}-darwin main.go

run:
	./${BINARY_NAME}-linux

run_windows:
	./${BINARY_NAME}-windows

run_mac:
	./${BINARY_NAME}-mac

build_and_run:clean build run

clean:
	rm -f -r ./${BINARY_FILE_PATH}
	go clean
