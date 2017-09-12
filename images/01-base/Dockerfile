FROM coldewey/os-rootfs
RUN ln -s /dev/null /etc/udev/rules.d/80-net-name-slot.rules
# Cleanup Buildroot
RUN rm /sbin/poweroff /sbin/reboot /sbin/halt && \
    sed -i '/^root/s!/bin/sh!/bin/bash!' /etc/passwd && \
    echo "$DISTRIB_ID \n \l" > /etc/issue && \
    rm -rf /run \
       /linuxrc \
       /etc/os-release \
       /var/cache \
       /var/lock \
       /var/log \
       /var/run \
       /var/spool \
       /var/lib/misc && \
    mkdir -p \
       /home \
       /run \
       /var/cache \
       /var/lock \
       /var/log \
       /var/run \
       /var/spool \
       /tmp/.X11-unix && \
    passwd -l root && \
    addgroup -g 1100 rancher && \
    addgroup -g 1101 docker && \
    addgroup -g 1103 sudo && \
    adduser -u 1100 -G rancher -D -h /home/rancher -s /bin/bash rancher && \
    adduser -u 1101 -G docker -D -h /home/docker -s /bin/bash docker && \
    adduser rancher docker && \
    adduser rancher sudo && \
    adduser docker sudo && \
    echo '%sudo ALL=(ALL) ALL' >> /etc/sudoers
COPY inputrc /etc/inputrc
COPY growpart /usr/bin/growpart
COPY start_ntp.sh /bin/start_ntp.sh
COPY create_xauth.sh /bin/create_xauth.sh
RUN sed -i s/"partx --update \"\$part\" \"\$dev\""/"partx --update --nr \"\$part\" \"\$dev\""/g /usr/bin/growpart && \
    sed -i -e 's/duid/clientid/g' /etc/dhcpcd.conf && \
    sed -i 1,10d /etc/rsyslog.conf && \
    echo "*.*                /var/log/syslog" >> /etc/rsyslog.conf

ENV XSERVER_AUTHFILE="/var/run/xauth"
RUN echo 'ACTION=="add",' \
    'ATTRS{bInterfaceProtocol}=="02",' \
    'ATTRS{bInterfaceClass}=="03",' \
    'ATTRS{bInterfaceSubClass}=="01",' \
    'ENV{REMOVE_CMD}="system-docker exec xorg sh -C '"'XAUTHORITY=$XSERVER_AUTHFILE DISPLAY=:0 /usr/bin/xinput -cursor [transparent cursor]'"'",' \
    'RUN+="system-docker exec xorg sh -C '"'XAUTHORITY=$XSERVER_AUTHFILE DISPLAY=:0 /usr/bin/xinput -cursor [normal cursor]'"'"' > /etc/udev/rules.d/99-touch-cursor.rules
#RUN touch /etc/machine-id
# dump kernel log to console (but after we've finished booting)
#    echo "kern.*                /dev/console" >> /etc/rsyslog.conf

ENTRYPOINT ["/usr/bin/ros", "entrypoint"]
