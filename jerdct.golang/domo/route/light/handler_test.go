package light_test

import (
	"testing"
	"jerdct.golang/domo/route/light"
)

func TestPattern(t *testing.T) {
	ctrl := new(light.Controller)
	if pt := ctrl.Pattern(); pt != "/light" {
		t.Errorf("simple light pattern test failed: got %q", pt)
	}
}
