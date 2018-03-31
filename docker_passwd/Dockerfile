FROM ubuntu:16.04
RUN apt-get  update
RUN apt-get  upgrade -y
RUN apt-get  dist-upgrade
RUN apt-get install -y openssh-server \
    gcc \
    sudo \
    curl \
    vim \
    wget \
    net-tools \
    iproute2 \
    iputils-ping
RUN apt-get clean && rm -rf /var/lib/apt/lists/*
RUN mkdir /var/run/sshd
RUN sed -i 's/Port 22/Port 2222/g'  /etc/ssh/sshd_config
RUN sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
RUN sed -i 's/PubkeyAuthentication yes/PubkeyAuthentication no/g' /etc/ssh/sshd_config
RUN echo 'root:root' | chpasswd
RUN echo "nameserver 8.8.8.8" >> /etc/resolv.conf
# SSH login fix. Otherwise user is kicked off after login
RUN sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd
RUN apt-get update  -y
ENV NOTVISIBLE "in users profile"
RUN echo "export VISIBLE=now" >> /etc/profile
EXPOSE 2222
CMD ["/usr/sbin/sshd", "-D"]
