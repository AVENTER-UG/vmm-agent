#!/sbin/openrc-run

name=$RC_SVCNAME
description="vmm-agent"
supervisor="supervise-daemon"
command="/usr/local/bin/vmm-agent"
pidfile="/run/agent.pid"
command_user="vmm:vmm"

depend() {
	after net
}
