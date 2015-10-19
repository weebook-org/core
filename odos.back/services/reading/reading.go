package reading

import (
	"net/http"
	"odos.back/services/rest"
)

type ReadingService struct {
}

type ReadingEndPoint struct {
}

func (s *ReadingService) EndPoints() (e []rest.EndPoint) {
	return
}

func (r *ReadingEndPoint) Pattern() (s string) {
	return
}

func (r *ReadingEndPoint) ServeHTTP(http.ResponseWriter, *http.Request) {

}
