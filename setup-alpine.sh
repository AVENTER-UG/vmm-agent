#!/bin/sh

set -xe

apk add openrc
apk add util-linux

ln -s agetty /etc/init.d/agetty.ttyS0
echo ttyS0 > /etc/securetty
rc-update add agetty.ttyS0 default

echo "root:root"|chpasswd

echo "nameserver 1.1.1.1" >> /etc/resolv.conf

rc-update add devfs boot
rc-update add procfs boot
rc-update add sysfs boot

rc-update add agent boot

for d in bin etc lib root sbin usr; do tar c "/$d" | tar x -C /my-rootfs;done
for dir in dev proc run sys var; do mkdir /my-rootfs/${dir}; done