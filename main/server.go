package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(response http.ResponseWriter, request *http.Request) {
	time.Sleep(time.Second * 5)
	fmt.Fprintf(response, "Hello, %s\n", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8001", nil)
}
