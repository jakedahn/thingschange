Vagrant.configure('2') do |config|
  vm_ram = ENV['VAGRANT_VM_RAM'] || 2048
  vm_cpu = ENV['VAGRANT_VM_CPU'] || 2

  config.vm.box = "puphpet/ubuntu1404-x64"

  config.vm.network "private_network", ip: "3.3.3.3"

  config.vm.provider :virtualbox do |vb|
    vb.customize ["modifyvm", :id, "--memory", vm_ram, "--cpus", vm_cpu]
  end
  
  config.vm.provider "vmware_fusion" do |v|
    v.vmx["memsize"] = vm_ram
    v.vmx["numvcpus"] = vm_cpu
  end

  config.vm.network :forwarded_port, guest: 5000, host: 5000
  config.vm.network :forwarded_port, guest: 5672, host: 5672

  config.vm.provision :shell, :inline => "/vagrant/init.sh"
end
