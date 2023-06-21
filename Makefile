#Dockerfile vars

#vars
IMAGENAME=vmm-agent
REPO=localhost:5000
TAG=$(shell git describe)
BUILDDATE=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)
BRANCH=$(shell git symbolic-ref --short HEAD)

.PHONY: help build-bin kernel rootfs run

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
	@cd src; CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w -X main.BuildVersion=${BUILDDATE} -X main.GitVersion=${TAG} -extldflags \"-static\"" -o ../build/vmm-agent .

run:
	@echo ">>>> Run"
	@cd src; CGO_ENABLED=0 GOOS=linux go run .

rootfs:
	@echo ">>>> Build Rootfs"
	@cd build; ../resources/scripts/build-rootfs.sh	

start-vm:
	@echo ">>>> Start VM"
	@cd ./resources/scripts; sudo ./start-microvm.sh

stop-vm:
	@echo ">>>> Stop VM"
	@cd ./resources/scripts; sudo ./stop-microvm.sh

all: build-bin rootfs
