package submission

import (
	"github.com/bmizerany/pat"
	"net/http"
)

const (
	submissionInfo string = "/submissions/infos"
	subGetOne      string = "/submissions/:id"
)

func BindSubmission(p *pat.PatternServeMux) {
	p.Get(submissionInfo, http.HandlerFunc(infoSubmission))
	p.Get(subGetOne, http.HandlerFunc(getOneSub))
}

func infoSubmission(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world submission !"))
}

func getOneSub(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	id := v.Get(":id")
	w.Write([]byte("Hello " + id))
}
