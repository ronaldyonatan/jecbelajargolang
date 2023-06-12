package main

import (
	"testing"
)

func TestPertama(t *testing.T) {
	result := iniLoopingWitParameter(3)

	if result != "fizz" {
		t.Error("value must be fizz, real : ", result)
	}
}
