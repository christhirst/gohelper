package iasserts

import (
	"reflect"
	"testing"
)

func AssertNotComparable[T any](t *testing.T, got, want T) {
	t.Helper()
	if equal := reflect.DeepEqual(got, want); !equal {
		t.Errorf("Want: %v is not equal to got: %v, equal: %v", want, got, equal)
	}

}

func AssertComparable[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("Want: %v is not equal to got: %v", want, got)
	}

}
