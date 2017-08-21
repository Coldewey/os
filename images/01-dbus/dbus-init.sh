#!/bin/ash
rm -f /var/run/dbus.pid
mkdir -p /var/run/dbus /run/lock

dbus-cleanup-sockets

exec dbus-daemon --system --nofork
