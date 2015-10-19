package services

import "net/http"

type ReadingService struct {
}

type ReadingEndPoint struct {
}

func (s *ReadingService) EndPoints() (e []EndPoint) {
	return
}

func (r *ReadingEndPoint) Pattern() (s string) {
	return
}

func (r *ReadingEndPoint) ServeHTTP(http.ResponseWriter, *http.Request) {

}
