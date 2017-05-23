rabbit_consumer
===============

Go program to sit and wait for messages

## Build
```
go build -v
```

## Usage
```
rabbit_consumer
```

Commandline parameters
* rabbithost, string, default:localhost
* rabbitport, int, default:5672

Environment variables
* RABBIT_HOSTNAME, overrides the rabbit host name, superceeds the rabbithost flag
* RABBIT_PORT, overrides the rabbit port, superceeds the rabbitport flag
