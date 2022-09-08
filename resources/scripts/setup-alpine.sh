#!/bin/sh

set -xe

apk add --no-cache openrc util-linux
apk add --no-cache gcc libc-dev
apk add --no-cache python3 go
apk add --no-cache g++ git cargo

mkdir /my-rootfs
mount /data/rootfs.ext4 /my-rootfs

cd /tmp
git clone --recurse-submodules https://github.com/bytecodealliance/wasmtime.git
cd wasmtime
cargo build --release
cp ./target/release/wasmtime /usr/local/bin/

apk del git cargo

ln -s agetty /etc/init.d/agetty.ttyS0
echo ttyS0 >/etc/securetty
rc-update add agetty.ttyS0 default

echo "root:root" | chpasswd

echo "nameserver 1.1.1.1" >>/etc/resolv.conf

addgroup -g 1000 -S vmm && adduser -u 1000 -S vmm -G vmm

rc-update add devfs boot
rc-update add procfs boot
rc-update add sysfs boot

rc-update add vmm-agent boot

for d in bin etc lib root sbin usr; do tar c "/$d" | tar x -C /my-rootfs; done
for dir in dev proc run sys var tmp; do mkdir /my-rootfs/${dir}; done

chmod 1777 /my-rootfs/tmp
mkdir -p /my-rootfs/home/vmm/
chown 1000:1000 /my-rootfs/home/vmm/

umount /my-rootfs

