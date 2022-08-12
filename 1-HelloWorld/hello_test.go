package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World!"

	if got != want {
		t.Errorf("Got %q but wanted %q", got, want)
	}
}

//To run the test you need to run "go mod init hello" first.
