package submission
import "net/http"


type reading struct {
}

func (v *reading) Pattern() string {
	return
}

func (v *reading) ServeHTTP(http.ResponseWriter, *http.Request) {

}
