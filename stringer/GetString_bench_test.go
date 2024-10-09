package stringer

import (
	"testing"
)

func BenchmarkStringerGetString(b *testing.B) {
	unit := NewGetStringerService()
	for range b.N {
		message := unit.GetString()
		print(message)
	}
}
