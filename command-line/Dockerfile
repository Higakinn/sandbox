FROM alpine

RUN apk add apg bash bash-completion \
            colordiff \
	    curl \ 
            coreutils \
            findutils \
            git \
            jq npm make tmux vim py3-pip

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

# install kubectl
ARG KUBECTL_VERSION=1.17.5
RUN curl -sLO https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl && \
    mv kubectl /usr/bin/kubectl && \
    chmod +x /usr/bin/kubectl

COPY .vimrc /root/.vimrc 
RUN curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
RUN vim -c PlugInstall -c q -c q

# install carbon-now-cli
#RUN npm i -g carbon-now-cli

WORKDIR /mnt
COPY .bashrc /root/.bashrc
