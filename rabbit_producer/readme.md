rabbit_producer
===============

Go program to send messages to a message queue

## Build
```
go build -v
```

## Usage

Run the rabbit consumer, then run a rabbit_producer

```
rabbit_producer
```

Send messages to /send URL in with curl (or wget)
```
curl -G -v "http://localhost:34500/send" --data-urlencode "message=hello world"
```

Commandline parameters
* port, int, string, default 34500, web server port to listen on
* rabbithost, string, default:localhost
* rabbitport, int, default:5672

Environment variables
* RABBIT_HOSTNAME, overrides the rabbit host name, superceeds the rabbithost flag
* RABBIT_PORT, overrides the rabbit port, superceeds the rabbitport flag
