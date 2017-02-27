package greeting

import "testing"

func TestGreeting(t *testing.T) {
	expected := "Hello, World!"
  observed := Greeting()
	if observed != expected {
		t.Fatalf("HelloWorld() = %v, want %v", observed, expected)
	}
}
