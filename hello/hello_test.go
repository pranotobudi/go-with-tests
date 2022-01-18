package hello

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	// Arrange
	// Act
	got := Hello()
	want := "Hello, world"

	// Assert
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")
	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
