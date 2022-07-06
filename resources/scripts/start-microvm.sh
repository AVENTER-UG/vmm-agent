#!/bin/bash

rm /tmp/firecracker.socket
firecracker  --api-sock /tmp/firecracker.socket --config-file ./alpine-config.json

