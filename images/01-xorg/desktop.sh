#!/bin/ash

rm -f /tmp/.X0-lock
rm -f /var/lib/xauth

mcookie=`/usr/bin/mcookie`
if test x"$mcookie" = x; then
   echo "Couldn't create cookie"
   exit 1
fi

xserverauthfile="/var/run/xauth"

xauth -q -f "$xserverauthfile" << EOF
add :0 . $mcookie
EOF

exec /usr/bin/Xorg :0 -auth "$xserverauthfile" -nolisten tcp $DISP_VT
