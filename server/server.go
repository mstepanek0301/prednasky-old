package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	http.Handle("/prednasky/",http.StripPrefix(
		"/prednasky/", http.FileServer(http.Dir(".")),
	))

	port := flag.String("p", "8000", "port to serve on")
	flag.Parse()
	log.Printf("Serving on HTTP port: %s\n", *port)
	log.Fatal(http.ListenAndServe(":" + *port, nil))
}
