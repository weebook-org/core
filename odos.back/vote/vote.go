package vote

import (
	"net/http"
	"github.com/bmizerany/pat"
)

const (
	voteInfo string = "/votes/infos"
)

func BindVote(m *pat.PatternServeMux) {
	m.Get(voteInfo, http.HandlerFunc(infoVote))
}

func infoVote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world vote !"))
}
