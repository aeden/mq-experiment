// A simple web app
package main

import (
	"fmt"
	"github.com/nats-io/nats"
	"html"
	"log"
	"net"
	"net/http"
	"os"
)

var version string

var (
	httpBindAddress = os.Getenv("HTTP_BIND_ADDRESS")
	httpBindPort    = os.Getenv("HTTP_BIND_PORT")
)

func main() {
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/version", versionHandler)
	httpMux.HandleFunc("/ping", pingHandler)
	//httpMux.HandleFunc("/", rootHandler)

	httpHostAndPort := net.JoinHostPort(httpBindAddress, httpBindPort)
	log.Print(fmt.Sprint("Starting HTTP service on ", httpHostAndPort))

	httpSrv := &http.Server{
		Addr:    httpHostAndPort,
		Handler: httpMux,
	}

	err := httpSrv.ListenAndServe()
	if err != nil {
		log.Fatal("Cannot start HTTP service", err)
	}
}

func handleError(w http.ResponseWriter, r *http.Request, err error, msg string, httpCode int) {
	log.Print(msg)
	http.Error(w, msg, httpCode)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "alive: %q", html.EscapeString(r.URL.Path))
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "version: %v", version)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		handleError(w, r, err, fmt.Sprintf("Failed to connect to message queue: %s", err), http.StatusInternalServerError)
		return
	}

	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		handleError(w, r, err, fmt.Sprintf("Failed to create encoded connection for message queue: %s", err), http.StatusInternalServerError)
		return
	}

	defer c.Close()
	c.Publish("ping", "test")

	fmt.Fprintf(w, "ping sent")
}
