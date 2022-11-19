FROM mysql:5.7

RUN apt update && apt install -y python-pip git vim curl

RUN pip install --upgrade sqlparse
RUN curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

COPY .vimrc /root/.vimrc
RUN vim -c PlugInstall -c q -c q
