package integers

import "testing"

func TestAdder(t *testing.T) {
	got := Adder(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
