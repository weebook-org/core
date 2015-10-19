package submission
import "net/http"


type saving struct {
}

func (v *saving) Pattern() string {
	return
}

func (v *saving) ServeHTTP(http.ResponseWriter, *http.Request) {

}
