package submission

import (
	"github.com/bmizerany/pat"
	"net/http"
)

const (
	commentsInfo string = "/submissions/comments/infos"
)

func BindComments(p *pat.PatternServeMux) {
	p.Get(commentsInfo, http.HandlerFunc(infoComments))
}

func infoComments(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world comments !"))
}
