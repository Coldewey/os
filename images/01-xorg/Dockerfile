FROM coldewey/os-base
ENV DISP_VT=vt2

ADD desktop.sh /usr/bin/display

COPY build/lsb-release /etc/
COPY prompt.sh /etc/profile.d/
