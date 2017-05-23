// https://golang.org/doc/articles/wiki/
package main

import (
	"flag"
	"fmt"
)

var (
	rabbitHostPtr = flag.String("rabbithost", "localhost", "Name of the rabbit host")
	rabbitPortPtr = flag.Int("rabbitport", 5672, "Port on which rabbit is running on")
)

func main() {
	flag.Parse()
	rabbitURI := rabbitAMQP()
	println(fmt.Sprintf("Rabbit messages being sent to %s", rabbitURI))
	receiveMessage()
}
