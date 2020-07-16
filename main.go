package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", landingPage)
	http.ListenAndServe(":8080", nil)
}

func landingPage(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Welcome to the landing page!")
}
