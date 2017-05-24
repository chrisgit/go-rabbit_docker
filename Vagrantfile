VAGRANTFILE_API_VERSION = '2'

$create_go_folders = <<EOF
mkdir -p /opt/go/bin
mkdir -p /opt/go/pkg
EOF

$build_rabbit_producer = <<EOF
cd /opt/go/src
export GOPATH=/opt/go
docker run --rm --name go-build -v /opt/go/src/rabbit_producer:/usr/src/rabbit_producer -v "$GOPATH":/go -w /usr/src/rabbit_producer golang sh -c " go get github.com/streadway/amqp  ; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo"
EOF

$build_rabbit_consumer = <<EOF
cd /opt/go/src
export GOPATH=/opt/go
docker run --rm --name go-build -v /opt/go/src/rabbit_consumer:/usr/src/rabbit_consumer -v "$GOPATH":/go -w /usr/src/rabbit_consumer golang sh -c " go get github.com/streadway/amqp  ; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo"
EOF

$build_rabbit_producer_docker_image = <<EOF
cd /opt/go/src/rabbit_producer
docker build . -t rabbit_producer
EOF

$build_rabbit_consumer_docker_image = <<EOF
cd /opt/go/src/rabbit_consumer
docker build . -t rabbit_consumer
EOF

$run_rabbit_container = <<EOF
docker run -d --hostname go-rabbit --name rabbitmq -p 5672:5672 -p 8080:15672 rabbitmq:3.6.9-management
EOF

$run_producer_consumer_containers = <<EOF
docker run -d --link rabbitmq:rabbitmq -p 34500:34500  --name rabbit_producer -e RABBIT_HOSTNAME=rabbitmq rabbit_producer
sleep 5
docker run -d --link rabbitmq:rabbitmq --name rabbit_consumer -e RABBIT_HOSTNAME=rabbitmq rabbit_consumer
EOF

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = 'ubuntu/trusty64'

  config.vm.define 'use-rabbitmq' do |v|
    v.vm.hostname = 'rabbitmq-container'
    v.vm.provider 'virtualbox' do |vb|
      vb.customize ['setextradata', 'global', 'GUI/SuppressMessages', 'all' ]
      vb.customize ['modifyvm', :id, '--ioapic', 'on']
      vb.cpus = 1
      vb.memory = 1024
      vb.linked_clone = true
    end
    v.vm.network 'private_network', ip: '172.20.20.10'
    v.vm.provision "docker" do |d|
      d.pull_images 'rabbitmq:3.6.9-management'
	  d.pull_images 'golang'
    end
    v.vm.network "forwarded_port", id: "rabbit_management", host: 8080, guest: 8080, protocol: "tcp"
    v.vm.synced_folder ".", "/opt/go/src", create: true
    v.vm.provision "shell", inline: $create_go_folders
    v.vm.provision "shell", inline: $build_rabbit_producer
    v.vm.provision "shell", inline: $build_rabbit_consumer
    v.vm.provision "shell", inline: $build_rabbit_producer_docker_image
    v.vm.provision "shell", inline: $build_rabbit_consumer_docker_image	
    v.vm.provision "shell", inline: $run_rabbit_container
    v.vm.provision "shell", inline: $run_producer_consumer_containers
  end
end
