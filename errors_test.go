package errors

import "testing"

func TestWrap(t *testing.T) {
	err := Wrap(nil, "text")
	if err != nil {
		t.Fatal(err)
	}
}

func TestWrapf(t *testing.T) {
	err := Wrapf(nil, "%s", "text")
	if err != nil {
		t.Fatal(err)
	}
}
