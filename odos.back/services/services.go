package services

import "net/http"

type Service interface {
	EndPoints() []EndPoint
}

type EndPoint interface {
	Pattern() string
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func BindServices() {

}
