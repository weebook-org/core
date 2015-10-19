package submission_test

import (
	"odos.back/services/submission"
	"testing"
)

func TestCommentsPath(t *testing.T) {
	c := submission.NewComment()
	if c.Pattern() != "/submissions/comments" {
		t.Error("bad path for comments end point")
	}
}
