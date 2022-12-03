FROM alpine

RUN apk add bash bash-completion colordiff jq curl py3-pip npm coreutils git make

# install csvkit
# https://csvkit.readthedocs.io/en/latest/
RUN pip install csvkit
RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/main' >> /etc/apk/repositories
RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/community' >> /etc/apk/repositories
RUN apk update
RUN apk add mongodb=3.4.4-r0

# install fpp
RUN cd /usr/local/ && \
    git clone https://github.com/facebook/PathPicker.git && \
    cd PathPicker/ && \
    ln -s "$(pwd)/fpp" /usr/local/bin/fpp

# install carbon-now-cli
#RUN npm i -g carbon-now-cli

WORKDIR /mnt

COPY .bashrc /root/.bashrc
