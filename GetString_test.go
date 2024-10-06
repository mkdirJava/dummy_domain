package getstring

import "testing"

func Test_getString(t *testing.T) {
	unit := NewGetStringerService()
	result := unit.GetString()
	if result != "hi" {
		t.Errorf("The GetString should return hi it got %s", result)
	}
}
