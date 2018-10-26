FROM golang:1.11.0-stretch AS base

# build the AuthorizedKeysCommand
RUN go get github.com/TACC-Cloud/ssh-keys-client/authorizedkeycommand



FROM centos:centos7

# Add default user.
ARG USER=docker
ARG USERHOME=/home/${USER}
ARG SSHDIR=${USERHOME}/.ssh

# Install openssh.
# Add host keys.
# Specify location of login banner.
RUN yum -y update && yum -y install openssh-server passwd sudo && yum clean all \
    && mkdir /var/run/sshd \
    && cd /etc/ssh && ssh-keygen -A -N '' \
    && sed -i "s/#Banner.*/Banner \/etc\/mybanner/g" /etc/ssh/sshd_config \
    && sed -i "s/#LogLevel .*/LogLevel DEBUG/g" /etc/ssh/sshd_config \
    && sed -i "s/#AuthorizedKeysCommand .*/AuthorizedKeysCommand \/usr\/local\/bin\/userkeys/g" /etc/ssh/sshd_config \
    && sed -i "s/#AuthorizedKeysCommandUser .*/AuthorizedKeysCommandUser nobody/g" /etc/ssh/sshd_config

# Set a welcome message for when a user sshs into the container.
ADD welcome_msg.txt /etc/mybanner

# Add new user.
# Set user's password as the name of the user.
# Add user to sudoers with no password.
# Set permissions for ~/.ssh and contents.
ADD ssh/ ${SSHDIR}/
RUN adduser ${USER} \
    && echo -e "${USER}\n${USER}" | (passwd --stdin ${USER}) \
    && echo "${USER}   ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers \
    && chmod -R 600 ${SSHDIR}/* \
    && chown -R ${USER}:${USER} ${SSHDIR}

WORKDIR /app
COPY --from=base /go/bin/authorizedkeycommand /usr/local/bin/userkeys

ENV HOME=/home/${USER}
RUN chmod 755 /usr/local/bin/userkeys

# Run ssh daemon and keys service.
EXPOSE 22
ADD run-keys-service.sh .
ENTRYPOINT ["./run-keys-service.sh"]
