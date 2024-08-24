package main

import "testing"

func TestSayHello(t *testing.T) {
	got := SayHi()
	want := "Salam"

	if got != want {
		t.Error("got %q want %q, got, other thing ")
	}
}
