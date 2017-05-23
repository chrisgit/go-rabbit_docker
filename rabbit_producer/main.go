// https://golang.org/doc/articles/wiki/
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var (
	portPtr       = flag.Int("port", 34500, "WebServer port to run the server on")
	rabbitHostPtr = flag.String("rabbithost", "localhost", "Name of the rabbit host")
	rabbitPortPtr = flag.Int("rabbitport", 5672, "Port on which rabbit is running on")
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Please send a message using /send?message=some text")
}

func messageSend(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	sendMessage(message)
	fmt.Fprintf(w, "Sent message %s", message)
}

func webServerPort() int {
	port := os.Getenv("SIMPLE_WEB_PORT")
	if port != "" {
		portNumber, _ := strconv.Atoi(port)
		return portNumber
	}
	return *portPtr
}

func main() {
	flag.Parse()
	serverPort := fmt.Sprintf(":%d", webServerPort())
	rabbitURI := rabbitAMQP()
	println(fmt.Sprintf("Web server running on port %s", serverPort))
	println(fmt.Sprintf("Rabbit messages being sent to %s", rabbitURI))
	http.HandleFunc("/", root)
	http.HandleFunc("/send", messageSend)
	http.ListenAndServe(serverPort, nil)
}
