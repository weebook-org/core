package submission

import (
	"odos.back/services/rest"
)

type Service struct {
	enpoints []rest.EndPoint
}

func NewService() (s *Service) {
	s = new(Service)
	s.enpoints = []rest.EndPoint{NewSubmission(), NewComment()}
	return
}

func (s *Service) EndPoints() []rest.EndPoint {
	return s.enpoints
}
