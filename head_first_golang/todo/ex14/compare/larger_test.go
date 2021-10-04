package compare

import "testing"

func TestFirstLarger(t *testing.T) {
	want := 2
	got := Larger(2, 3)
	if got != want {
		t.Errorf("Larger(%d, %d) = %d want %d", 2, 3, got, want)
	}
}

func TestTwoLarger(t *testing.T) {
	want := 1
	got := Larger(2, 1)
	if got != want {
		t.Errorf("Larger(%d, %d) = %d want %d", 2, 1, got, want)
	}
}
