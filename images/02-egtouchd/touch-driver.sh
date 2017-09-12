#!/bin/sh
COOKIE=`xauth list | grep $DISPLAY | head -n 1 | awk -F ' ' '{print $3}'`

xauth add $DISPLAY . $COOKIE

if [[ -n `lsusb | grep 0eef` ]]; then
    modprobe -r usbtouchscreen

    touch /tmp/eGTouch_tmp
    eGTouchD

    tail -f $(ls -1 /tmp/eGTouch_* )
else
    exec true
fi
