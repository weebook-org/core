package submission_test

import (
	"odos.back/services/submission"
	"testing"
)

func TestSubmissionPath(t *testing.T) {
	s := submission.NewSubmission()
	if s.Pattern() != "/submissions" {
		t.Error("bad path for submission end point")
	}
}
