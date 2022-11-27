FROM alpine

RUN apk add bash colordiff jq curl py3-pip npm 

# install csvkit
# https://csvkit.readthedocs.io/en/latest/
RUN pip install csvkit

# install carbon-now-cli
RUN npm i -g carbon-now-cli

WORKDIR /mnt

COPY .bashrc /root/.bashrc
