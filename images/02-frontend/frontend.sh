#!/bin/sh
COOKIE=`xauth list | grep $DISPLAY | head -n 1 | awk -F ' ' '{print $3}'`

xauth add $DISPLAY . $COOKIE
electron
