package main

import (
	"net/http"
	"log"
	"jerdct.golang/domo/route"
)

func main() {
	route.Build()
	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		log.Printf("The server stop because of %v", error)
	}
}
