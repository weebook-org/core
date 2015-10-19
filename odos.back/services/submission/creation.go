package submission
import "net/http"


type creation struct {
}

func (v *creation) Pattern() string {
	return
}

func (v *creation) ServeHTTP(http.ResponseWriter, *http.Request) {

}