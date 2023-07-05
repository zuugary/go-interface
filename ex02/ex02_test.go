package main_test

import (
	"testing"
	"time"

	ex "github.com/zuugary/go-interface/ex02"
)

type fakeClock struct{}

func (f *fakeClock) Sleep(d time.Duration) {}

func TestLongProcessCalc(t *testing.T) {
	got := ex.LongProcessCalc(2, 3, &fakeClock{})
	want := 5

	if got != want {
		t.Errorf("got: %d; want: %d\n", got, want)
	}
}
