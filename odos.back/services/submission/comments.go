package submission

import "net/http"

const (
	commentsPath string = "/comments"
)

type Comments struct {
}

func (v *Comments) Pattern() string {
	return commentsPath
}

func (v *Comments) ServeHTTP(http.ResponseWriter, *http.Request) {

}
