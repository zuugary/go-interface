package main_test

import (
	"testing"

	ex "github.com/zuugary/go-interface/ex01"
)

func TestLongProcessCalc(t *testing.T) {
	got := ex.LongProcessCalc(2, 3)
	want := 5

	if got != want {
		t.Errorf("got: %d; want: %d\n", got, want)
	}
}
