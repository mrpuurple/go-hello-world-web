package main

import (
	"fmt"
	"net/http"
	"github.com/mrpuurple/go-hello-world-web/pkg/render"
)

const portNumber = "8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting webserver on port %q\n", portNumber)
	_ = http.ListenAndServe(":" + portNumber, nil)
}
