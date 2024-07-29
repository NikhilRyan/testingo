package assertions

import (
	"strings"
	"testing"
)

func AssertEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func AssertNotEqual(t *testing.T, expected, actual interface{}) {
	if expected == actual {
		t.Errorf("Expected %v to be different from %v", expected, actual)
	}
}

func AssertNil(t *testing.T, actual interface{}) {
	if actual != nil {
		t.Errorf("Expected nil but got %v", actual)
	}
}

func AssertNotNil(t *testing.T, actual interface{}) {
	if actual == nil {
		t.Errorf("Expected not nil but got nil")
	}
}

func AssertTrue(t *testing.T, condition bool) {
	if !condition {
		t.Errorf("Expected condition to be true, but it was false")
	}
}

func AssertFalse(t *testing.T, condition bool) {
	if condition {
		t.Errorf("Expected condition to be false, but it was true")
	}
}

func AssertContains(t *testing.T, container, item string) {
	if !strings.Contains(container, item) {
		t.Errorf("Expected '%s' to contain '%s'", container, item)
	}
}
