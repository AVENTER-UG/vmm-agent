#Dockerfile vars

#vars
IMAGENAME=vmm-agent
REPO=localhost:5000
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
	@cd build; docker run -i -u 1000:1000 -v $PWD:/data avhost/ubuntu_build:jammy /bin/bash < ../resources/scripts/build-kernel.sh

rootfs:
	@echo ">>>> Build Rootfs"
	@cd build; sudo ../resources/scripts/build-rootfs.sh	

start-vm:
	@echo ">>>> Start VM"
	@cd ./resources/scripts; sudo ./start-microvm.sh

stop-vm:
	@echo ">>>> Stop VM"
	@cd ./resources/scripts; sudo ./stop-microvm.sh

all: build-bin kernel rootfs
