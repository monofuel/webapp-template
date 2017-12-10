# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|

  config.vm.box = "debian/jessie64"
  config.vm.network "forwarded_port", guest: 3002, host: 3002

  config.vm.provision "shell", inline: <<-SHELL
  cd /vagrant/mono-scripts/setup/
  export MONO_USER='vagrant'
  bash golang.bash
  bash nodejs.bash

  ln -s /home/vagrant/go/src/github.com/monofuel/webapp-template /vagrant
  cd /home/vagrant/go/src/github.com/monofuel/webapp-template
  make setup

  SHELL
end
