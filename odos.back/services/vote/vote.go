package services
import (
	"net/http"
	"odos.back/services"
)


type voteService struct {
}

type voteRegistration struct {
}

func (s *voteService) EndPoints() []services.EndPoint {
	return
}

func (v *voteRegistration) Pattern() string {
	return
}

func (v *voteRegistration) ServeHTTP(http.ResponseWriter, *http.Request) {

}
