package submission

import "net/http"

const (
	submissionPath string = "/submissions"
)

type Submission struct {
}

func NewSubmission() (s *Submission) {
	s = new(Submission)
	return
}

func (v *Submission) Pattern() string {
	return submissionPath
}

func (v *Submission) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
