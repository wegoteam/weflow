BINARY_NAME=weflow
BINARY_FILE_PATH=bin

build:
	mkdir -p ${BINARY_FILE_PATH}
	cp -r ./config ./${BINARY_FILE_PATH}/config
	GOARCH=amd64 GOOS=linux go build -o ./${BINARY_FILE_PATH}/${BINARY_NAME} main.go

build_windows:
	mkdir -p ${BINARY_FILE_PATH}
	cp -r ./config ./${BINARY_FILE_PATH}/config
	GOARCH=amd64 GOOS=windows go build -o ./${BINARY_FILE_PATH}/${BINARY_NAME}.exe main.go

build_mac:
	mkdir -p ${BINARY_FILE_PATH}
	cp -r ./config ./${BINARY_FILE_PATH}/config
	GOARCH=amd64 GOOS=darwin go build -o ./${BINARY_FILE_PATH}/${BINARY_NAME} main.go

run:
	./${BINARY_FILE_PATH}/${BINARY_NAME}

run_windows:
	./${BINARY_FILE_PATH}/${BINARY_NAME}

run_mac:
	./${BINARY_FILE_PATH}/${BINARY_NAME}

build_and_run:clean build run

clean:
	rm -f -r ./${BINARY_FILE_PATH}
	go clean
