package greeting

import "testing"

func TestGreeting(t *testing.T) {
	expected := "Hello, World!"
	observed := Greeting()
	if observed != expected {
		t.Fatalf("observed %v, expected %v", observed, expected)
	}
}
