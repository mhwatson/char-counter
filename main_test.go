package main

import "testing"

func TestTruth(t *testing.T) {
	if true != true {
		t.Error("Expected true to be true")
	}
}
