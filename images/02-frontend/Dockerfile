FROM coldewey/os-base

ENV DISPLAY=:0
ENV XAUTHORITY=/var/run/xauth

#RUN adduser -D -u 1100 rancher
#RUN adduser -D -u 1101 docker

#RUN sed -i 's/rancher:!/rancher:*/g' /etc/shadow && \
#    sed -i 's/docker:!/docker:*/g' /etc/shadow && \
#    echo '## allow password less for rancher user' >> /etc/sudoers && \
#    echo 'rancher ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers && \
#    echo '## allow password less for docker user' >> /etc/sudoers && \
#    echo 'docker ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers

COPY prompt.sh /etc/profile.d/
COPY frontend.sh /usr/bin/frontend

RUN npm install -g electron
