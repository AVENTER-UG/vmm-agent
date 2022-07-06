#!/bin/bash

set -xe

dd if=/dev/zero of=rootfs.ext4 bs=1M count=1000
mkfs.ext4 rootfs.ext4
mkdir -p /tmp/my-rootfs
mount rootfs.ext4 /tmp/my-rootfs

docker run -i --rm \
    -v /tmp/my-rootfs:/my-rootfs \
    -v "$(pwd)/vmm-agent:/usr/local/bin/vmm-agent" \
    -v "$(pwd)/../resources/scripts/openrc-service.sh:/etc/init.d/vmm-agent" \
    alpine sh < ../resources/scripts/setup-alpine.sh

umount /tmp/my-rootfs
