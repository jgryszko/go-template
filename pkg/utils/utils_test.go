package utils

import "testing"

// TestGreet tests the Greet function with different inputs
func TestGreet(t *testing.T) {
	expected := "Hello, Bill"
	result := Greet("Bill")

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	expected = "Hello, Alice"
	result = Greet("Alice")

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
