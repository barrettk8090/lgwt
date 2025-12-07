package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat the standard five times", func(t *testing.T) {
		got := Repeat("a", 5)
		expected := "aaaaa"

		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})
	t.Run("repeat a custom number of times", func(t *testing.T) {
		got := Repeat("a", 12)
		expected := "aaaaaaaaaaaa"

		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})
}

// TODO: assetCorrectMessage

// TODO: Handle negative numbers

// TODO: Handle uppercase letters

func ExampleRepeat() {
	iterationSample := Repeat("r", 7)
	fmt.Println(iterationSample)
	// Output: rrrrrrr
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 6)
	}
}
