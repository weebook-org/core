package submission_test

import (
	"odos.back/services/submission"
	"testing"
)

func TestSubmissionPath(t *testing.T) {
	s := new(submission.Submission)
	if s.Pattern() != "/submission" {
		t.Error("bad path for submission end point")
	}
}
