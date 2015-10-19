package submission

import "net/http"

const (
	commentsPath string = "/submissions/comments"
)

func NewComment() (c *Comments) {
	c = new(Comments)
	return
}

type Comments struct {
}

func (v *Comments) Pattern() string {
	return commentsPath
}

func (v *Comments) ServeHTTP(http.ResponseWriter, *http.Request) {

}
