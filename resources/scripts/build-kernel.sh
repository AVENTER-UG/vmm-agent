#!/bin/bash

set -xe
cd /tmp
git clone https://github.com/torvalds/linux.git linux
cd linux
git checkout v5.14
curl https://raw.githubusercontent.com/AVENTER-UG/vmm-agent/resources/kernel_config/microvm-kernel-x86_64-5.14.config > .config

make -i -j4 vmlinux
cp vmlinux /data

# uncompressed kernel image available under ./vmlinux
