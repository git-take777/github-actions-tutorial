
package main

import "testing"

func TestEvenOrOdd(t \*testing.T) {
result := EvenOrOdd(2)
if result != "even" {
t.Errorf("Expected 'even', but got '%s'", result)
}
}