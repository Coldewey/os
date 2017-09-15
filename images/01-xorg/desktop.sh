#!/bin/ash

rm -f /tmp/.X0-lock

if [[ -n `lsusb | grep 0eef` ]]; then
    NOCURSOR="-nocursor"
fi

exec /usr/bin/Xorg :0 -auth "$XSERVER_AUTHFILE" $NOCURSOR -nolisten tcp $DISP_VT
