package main

import "testing"

func TestCountDown(t *testing.T) {
	got := CountDown(5)
	want := "5\n4\n3\n2\n1"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
