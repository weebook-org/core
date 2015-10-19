package services
import "net/http"


type service interface {
	endPoints()[]endPoint
}

type endPoint interface {
	pattern() string
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func BindServices() {

}
