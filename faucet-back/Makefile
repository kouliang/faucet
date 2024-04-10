.PHONY: all help clean gotool build run

BINARY="faucet-app"

all: clean gotool build run

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	@if [ -f tap.log ] ; then rm tap.log ; fi

gotool:
	go fmt ./
	go vet ./

build:
	go build -o ./${BINARY}

buildlinux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./${BINARY}

run:
	./${BINARY} config.toml

dockerbuild:
	docker build . -t tap_app

dockerrun:
	docker run -p 9988:8090 tap_app 

help:
	@echo "make         - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make clean   - 移除二进制文件和 vim swap files"
	@echo "make gotool  - 运行 Go 工具 'fmt' and 'vet'"
	@echo "make build   - 编译 Go 代码, 生成二进制文件"
	@echo "make run     - 运行"

	@echo "make buildlinux   - 编译 Go 代码, 生成liunx平台二进制文件"
	@echo "make dockerbuild  - 编译为 docker 镜像"
	@echo "make dockerrun    - 运行 docker 镜像"
