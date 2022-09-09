#!/bin/bash

cd /tmp
git clone https://github.com/torvalds/linux.git linux
cd linux
git checkout v5.14
curl https://raw.githubusercontent.com/AVENTER-UG/vmm-agent/main/resources/kernel_config/microvm-kernel-x86_64-5.14.config > .config

make -i -j1 vmlinux
cp vmlinux /data
ls -l /data
df -h

# uncompressed kernel image available under ./vmlinux
