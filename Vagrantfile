# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "centos/7"

  config.vm.provision :docker

  config.vm.synced_folder "./", "/home/vagrant/shared"

  config.vm.provision "shell", inline: <<-SHELL
	# install docker-compose
	curl -L "https://github.com/docker/compose/releases/download/1.27.4/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
	chmod +x /usr/local/bin/docker-compose

	# install kubectl
	curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.26.0/bin/linux/amd64/kubectl
	chmod +x ./kubectl
	mv ./kubectl /usr/local/bin/kubectl

	# install packages
	yum -y install epel-release.noarch vim tmux tig

        # tmux setting
        cp /home/vagrant/shared/command-line/.tmux.conf ~/.tmux.conf

        # vim setting 
        cp /home/vagrant/shared/command-line/.vimrc ~/.vimrc
	curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
	vim -c PlugInstall -c q -c q
  SHELL
end
