package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	result := fibonacci(6)

	if result != 8 {
		t.Error("value must be 8, result : ", result)
	}
}
