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
	stop_weflow
	./${BINARY_FILE_PATH}/${BINARY_NAME}

run_windows:
	stop_weflow
	./${BINARY_FILE_PATH}/${BINARY_NAME}

run_mac:
	stop_weflow
	./${BINARY_FILE_PATH}/${BINARY_NAME}

stop_weflow:
	ps -ef | grep ${BINARY_NAME} | grep -v grep | awk '{print $2}' | xargs kill -9

build_and_run:clean build run

clean:
	rm -f -r ./${BINARY_FILE_PATH}
	go clean
