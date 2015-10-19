package submission_test

import (
	"odos.back/services/submission"
	"testing"
)

func TestCommentsPath(t *testing.T) {
	c := new(submission.Comments)
	if c.Pattern() != "/comments" {
		t.Error("bad path for comments end point")
	}
}
