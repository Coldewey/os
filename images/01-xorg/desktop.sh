#!/bin/ash

rm -f /tmp/.X0-lock

exec /usr/bin/Xorg :0 -auth "$XSERVER_AUTHFILE" -nolisten tcp $DISP_VT
