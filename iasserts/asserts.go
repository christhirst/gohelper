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

func AssertNoError(t testing.TB, got interface{}) {
	t.Helper()
	if got != nil {
		t.Errorf("expected no error got %v", got)
	}
}

func AssertError(t testing.TB, got interface{}) {
	t.Helper()
	if got == nil {
		t.Errorf("expected error got %v", got)
	}
}

func AssertEmptyString(t testing.TB, got string) {
	t.Helper()
	if got == "" {
		t.Errorf("expected string but got %v", got)
	}
}

func AssertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %v, expected: %v", got, want)
	}
}

func AssertResponseCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected  %v, got %v", want, got)
	}
}
func AssertCorrectMessage(t testing.TB, got, want map[string][]string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Not equal maps, got %q want %q", got, want)
	}
}
func AssertFormError(t testing.TB, got error) {
	t.Helper()
	if got == nil {
		t.Errorf("response body is wrong, got %q ", got)
	}
}

func AssertGeneric[k comparable](t testing.TB, got, want k) {
	t.Helper()
	if got != want {
		t.Errorf("got: %v, expected: %v", got, want)
	}
}
