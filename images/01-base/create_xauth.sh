#!/bin/sh
rm -f /var/lib/xauth

mcookie=`/usr/bin/mcookie`
if test x"$mcookie" = x; then
   echo "Couldn't create cookie"
   exit 1
fi

xauth -q -f "$XSERVER_AUTHFILE" << EOF
add :0 . $mcookie
EOF