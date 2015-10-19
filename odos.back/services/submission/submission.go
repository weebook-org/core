package submission

import "net/http"

const (
	submissionPath string = "/submission"
)

type Submission struct {
}

func (v *Submission) Pattern() string {
	return submissionPath
}

func (v *Submission) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
