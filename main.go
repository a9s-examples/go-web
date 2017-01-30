package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func helloWorldHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func main() {
	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "4000"
	}
	http.HandleFunc("/", helloWorldHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
