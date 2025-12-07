package iteration

import "testing"

func TestIterate(t *testing.T) {
	iteration := Repeat("a")
	expected := "aaaaa"

	if iteration != expected {
		t.Errorf("expected %q but got %q", expected, iteration)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
