package submission
import "net/http"


type edition struct {
}

func (v *edition) Pattern() string {
	return
}

func (v *edition) ServeHTTP(http.ResponseWriter, *http.Request) {

}
