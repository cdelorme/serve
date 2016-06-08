package main

import (
	"net/http"
	"os"

	"github.com/skratchdot/open-golang/open"
)

var run = open.Run

func main() {
	port := getPort()
	wd, _ := os.Getwd()
	go run("http://127.0.0.1:" + port)
	http.ListenAndServe(":"+port, http.FileServer(http.Dir(wd)))
}

func getPort() string {
	port := os.Getenv("PORT")
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	if len(port) == 0 {
		port = "3000"
	}
	return port
}
