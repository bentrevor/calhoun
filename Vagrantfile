# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.require_version ">= 1.7.4"

Vagrant.configure("2") do |config|

  config.vm.box = 'ubuntu/trusty64'
  config.vm.box_check_update = false
  config.vm.synced_folder '~/vms/calhoun', '/home/vagrant/go/src/github.com/bentrevor/calhoun', nfs: true
  config.vm.network :private_network, ip: '192.168.56.56'
  config.ssh.forward_agent = true

  config.vm.provision '~~ install git ~~', type: 'shell', inline: <<-SHELL
apt-get update
apt-get install -y git
SHELL

  install_golang = <<-SHELL
wget https://storage.googleapis.com/golang/go1.5.3.linux-amd64.tar.gz -q
tar -C /usr/local -xzf go1.5.3.linux-amd64.tar.gz
SHELL

  install_dev_env = <<-SHELL
git clone https://github.com/bentrevor/dev-env.git
/home/vagrant/dev-env/bootstrap --verbose --zsh --vim --dotfiles
SHELL

  provision_environment(config, 'dev', {
                          host_port: 4567,
                          memory: 4096,
                          cpus: 1,
                          scripts: {
                            'install golang' => install_golang,
                            'install dev-env' => install_dev_env
                          },
                        })

  provision_environment(config, 'prod', {
                          host_port: 7654,
                          memory: 16384,
                          cpus: 8,
                          scripts: {},
                        })
end

def provision_environment(config, env_name, opts)
  host_port = opts[:host_port]
  memory = opts[:memory]
  cpus = opts[:cpus]
  scripts = opts[:scripts]

  autostart = env_name == 'dev'

  config.vm.define env_name, autostart: autostart do |e|
    e.vm.network 'forwarded_port', host: host_port, guest: 8080, autocorrect: true
    e.vm.provider :virtualbox do |vb|
      vb.memory = memory
      vb.cpus = cpus
    end
  end

  scripts.each do |desc, script|
    config.vm.provision desc, type: 'shell', inline: script
  end
end
