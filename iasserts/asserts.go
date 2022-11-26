package iasserts

import (
	"reflect"
	"testing"
)

func AssertNotComparable[T any](t *testing.T, got, want T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Error()
	}

}

func AssertComparable[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Error()
	}

}
