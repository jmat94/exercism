package main

import "testing"

func TestHello(t *testing.T) {

	got := return_string()
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
