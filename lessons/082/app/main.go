package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"flag"
)

func main() {
	name := flag.String("name", "", "service name")
	flag.Parse()

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %s", r.Method)
		log.Printf("Protocol: %s", r.Proto)
		log.Printf("Headers: %v", r.Header)
		log.Printf("Body: %s", r.Body)
		log.Printf("RequestURI: %s", r.RequestURI)
		fmt.Fprintf(w, "Hello from %s, resource: %q\n", *name, html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
