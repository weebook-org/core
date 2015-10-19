package services
import "net/http"

type readingService struct {
}

type readingEndPoint struct {

}

func (s *readingService) endPoints()[]endPoint {
	return
}

func (r *readingEndPoint) pattern() string {
	return
}
func (r *readingEndPoint)ServeHTTP(http.ResponseWriter, *http.Request){

}
