FROM mysql:8.0-debian

RUN apt-get update && apt-get install -y python3-pip git vim curl

RUN pip3 install --upgrade sqlparse
RUN curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

COPY .vimrc /root/.vimrc
RUN vim -c PlugInstall -c q -c q
