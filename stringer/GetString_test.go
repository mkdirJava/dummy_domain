package stringer

import "testing"

// This setup and tear down effects tests that match on the
// command line
func TestMain(m *testing.M) {
	println("I am the setup")
	m.Run()
	println("I am the teardown")
}

// These tests run in parraelle
func Test_getString(t *testing.T) {
	unit := NewGetStringerService()

	t.Run("When there is allot of strings", func(t *testing.T) {
		result := unit.GetString()
		if result != "hi" {
			t.Errorf("The GetString should return hi it got %s", result)
		}
	})

	t.Run("When I want to go home", func(t *testing.T) {
		result := unit.GetString()
		if result != "hi" {
			t.Errorf("The GetString should return hi it got %s", result)
		}
	})
}
