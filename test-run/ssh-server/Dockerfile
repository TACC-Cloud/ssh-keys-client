FROM centos:centos7

# Add default user.
ARG USER=docker
ARG USERHOME=/home/${USER}
ARG SSHDIR=${USERHOME}/.ssh

# Add user and create a home directory.
# Add user to sudoers with no password.
# Set permissions for user's home directory.
# Unlock non-password user to enable SSH login.
# Create user's ssh dir.

# Install openssh.
# Add host keys.
# Specify location of login banner.
RUN yum -y update && yum -y install openssh-server passwd sudo && yum clean all \
    && mkdir /var/run/sshd \
    && cd /etc/ssh && ssh-keygen -A -N '' \
    && sed -i "s/#Banner.*/Banner \/etc\/mybanner/g" /etc/ssh/sshd_config

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

#USER $USER
WORKDIR /app

# Run ssh daemon and keys service.
EXPOSE 22
ADD run-keys-service.sh .
ENTRYPOINT ["./run-keys-service.sh"]
