package main
import (
	"odos.back/services"
	"net/http"
	"log"
)

func main() {
	services.BindServices()
	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		log.Printf("The server stop because of %v", error)
	}
}
