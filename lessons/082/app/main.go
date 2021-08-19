package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	// http.Handle("/foo", fooHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %s", r.Method)
		log.Printf("Protocol: %s", r.Proto)
		log.Printf("Headers: %v", r.Header)
		log.Printf("Host: %s", r.Host)
		log.Printf("Body: %s", r.Body)
		log.Printf("RequestURI: %s", r.RequestURI)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
