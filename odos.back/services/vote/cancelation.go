package vote
import "net/http"


type cancelation struct {
}

func (v *cancelation) Pattern() string {
	return
}

func (v *cancelation) ServeHTTP(http.ResponseWriter, *http.Request) {

}