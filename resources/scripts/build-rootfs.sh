#!/bin/bash

set -xe

dd if=/dev/zero of=rootfs.ext4 bs=1M count=1100
mkfs.ext4 rootfs.ext4

docker run -i --rm \
  -v "$(pwd):/data" \
    -v "$(pwd)/vmm-agent:/usr/local/bin/vmm-agent" \
    -v "$(pwd)/../resources/scripts/openrc-service.sh:/etc/init.d/vmm-agent" \
    -v "$(pwd)/../resources/scripts/extract-vmlinux:/usr/bin/extract-vmlinux" \
    --privileged \
    debian:11 sh < ../resources/scripts/setup-debian.sh

