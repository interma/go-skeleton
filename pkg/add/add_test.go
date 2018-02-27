package add

import "testing"

func TestAdd(t *testing.T) {
	t.Logf("in TestAdd")
	if Add(1, 2) != 3 {
		t.Errorf("test failed")
	}
}
