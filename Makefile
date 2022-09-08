#Dockerfile vars

#vars
IMAGENAME=vmm-agent
TAG=`git describe`
BUILDDATE=`date -u +%Y-%m-%dT%H:%M:%SZ`
BRANCH=`git symbolic-ref --short HEAD`

.PHONY: help build-bin kernel rootfs

help:
	    @echo "Makefile arguments:"
	    @echo ""
	    @echo "Makefile commands:"
			@echo "build-bin"
			@echo "kernel"
			@echo "rootfs"
			@echo ${TAG}

.DEFAULT_GOAL := all

build-bin:
	@echo ">>>> Build Binary"	
	@cd src; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.BuildVersion=${BUILDDATE} -X main.GitVersion=${TAG} -extldflags \"-static\"" -o ../build/vmm-agent main.go handle_wasm.go handle_python.go handle_cpp.go handle_golang.go handle_c.go exec.go

kernel:
	@echo ">>>> Build Kernel"	
	@curl -l https://dl-cdn.alpinelinux.org/alpine/v3.16/releases/x86_64/alpine-netboot-3.16.2-x86_64.tar.gz | tar -xzv boot/vmlinuz-lts  --strip-components=1 
	@mv vmlinuz-lts ./build/vmlinuz

rootfs:
	@echo ">>>> Build Rootfs"
	@cd build; ../resources/scripts/build-rootfs.sh	

start-vm:
	@echo ">>>> Start VM"
	@cd ./resources/scripts; sudo ./start-microvm.sh

stop-vm:
	@echo ">>>> Stop VM"
	@cd ./resources/scripts; sudo ./stop-microvm.sh

all: build-bin kernel rootfs
