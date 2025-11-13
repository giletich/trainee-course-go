package hello

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, world!"
	result := Hello()

	if result != expected {
		t.Errorf("Hello() = %q; expected %q", result, expected)
	}
}