# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|

  config.vm.box = 'centos/7'
  config.vm.synced_folder '.', '/vagrant', disabled: true
  config.vm.synced_folder ".", "/go/src/github.com/hajhatten/graphite-beacon-web", 
    type: "rsync",
    rsync__exclude: ".git/",
    rsync__args: ["--verbose", "--rsync-path='sudo rsync'", "--archive", "--delete", "-z"]

  
  config.vm.network "private_network", ip: "192.168.200.10"
  config.vm.network "forwarded_port", guest: 3000, host: 3000

  config.vm.provider 'virtualbox' do |vb|
    vb.name = 'vagrant-alpine-graphite-beacon-web'
    vb.cpus = 4
    vb.memory = 1024
    vb.gui = false
  end

end