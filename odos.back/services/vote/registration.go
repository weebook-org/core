package vote
import "net/http"

type registration struct {
}

func (v *registration) Pattern() string {
	return
}

func (v *registration) ServeHTTP(http.ResponseWriter, *http.Request) {

}
