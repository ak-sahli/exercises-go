package greeting

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
  observed := greeting()
	if observed != expected {
		t.Fatalf("HelloWorld() = %v, want %v", observed, expected)
	}
}
