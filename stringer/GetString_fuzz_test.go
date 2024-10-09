package stringer

import (
	"fmt"
	"testing"
)

func Fuzz_getStringer(f *testing.F) {
	for _, item := range []string{"someone", "bob", "Adele", "Tim"} {
		f.Add(item)
	}
	unit := NewGetStringerService()
	f.Fuzz(func(t *testing.T, arg string) {
		result := unit.Greeting(arg)
		expected := fmt.Sprintf("Hi there %s", arg)
		if result != expected {
			t.Fatalf("result should be '%s' instead got '%s'", expected, result)
		}
	})
}
