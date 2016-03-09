package main

import (
	"net/http"
	"os"
)

func main() {
	port := getPort()
	wd, _ := os.Getwd()
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
