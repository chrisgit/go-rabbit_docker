Go and Docker
=============

Rather than have a development environment on my laptop for ease of setup and speed I do the follow
* Run development environment in Virtual Machines
* Provision my development environment with Vagrant and/or Chef

Docker simplifies the process further, it has a pre-installed and configured environment ready to go.

As I'm experimenting with both Docker and Go I thought I'd combine the two, the initial goal was container to container communication.

## Vagrantfile
The power of this automation is in the Vagrantfile which includes comments.
The end result is a rabbitmq instance running and two compiled go programs
* rabbit_producer
* rabbit_consumer

## The code
The code has largely come from the tutorial at rabbit https://www.rabbitmq.com/tutorials/tutorial-one-go.html

rabbit_producer is a simple Go program that runs as a web service on a pre-defined port, it receives a message and passes it on to a rabbitmq queue
rabbit_consumer is a simple Go program that will receive a message from a rabbitmq queue

Rabbitmq is run in a container, the Rabbitmq image contains the management tools and is accessible on the docker host machine via port 8080; the vagrant file includes port forwarding rules that allow the management tools to be used by the host computer (i.e. your laptop or desktop, via http://localhost:8080).

## Compiling the code
The code is compiled using the official Golang docker image https://hub.docker.com/_/golang/

The compile instruction is via a simple go build command. Before compiling code the GOPATH needs to be setup and packages downloaded.

This is all handled by the Vagrantfile, nothing to do here but if you were to compile the code using normal go build it would be
```
cd /opt/go/src/rabbit_producer
sudo -E CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
```

sudo -E to pass in the GOPATH variable.

The build includes any libraries, this is because the Go program will be running a container built from Scratch, see https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/

## Running the code
After Vagrant has finished there should be a running instance of rabbitmq and four available Docker containers.

To keep things simple we will run the rabbit_producer and rabbit_consumer directly.

Login to the virtual machine with 

```
vagrant ssh
```

Change folders to /opt/go/src/rabbit_producer and run rabbit_producer, this sits and waits for messages on port 34500.

In another ssh session change folders to /opt/go/src/rabbit_consumer and run rabbit_consumer, this sits and waits for messages from rabbitmq.

In another ssh session run a curl command to send a message to the rabbit_producer.
```
curl -G -v "http://localhost:34500/send" --data-urlencode "message=hello world"
```

The message should be output by the rabbit_consumer.

## Running everything in containers

Alternatively you could run rabbit_producer and rabbit_consumer containers but first you would need the ip address of the rabbit container, which is obtained by running the following
```
docker inspect --format '{{ .NetworkSettings.IPAddress }}' rabbitmq
```

Run the producer against the IP address of the container (or use Docker Networking)
```
docker run -it -p 34500:34500 -e "RABBIT_HOSTNAME=172.17.0.2"  rabbit_producer
```

License and Authors
-------------------
Please see [LICENSE][licence]
Authors: Chris Sullivan

[licence]: https://github.com/chrisgit/go-rabbit_docker/blob/master/LICENSE