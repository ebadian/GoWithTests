package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("Got %q but wanted %q", got, want)
	}
}

//To run the test you need to run "go mod init hello" first.
//Writing a test is just like writing a function, with a few rules
//It needs to be in a file with a name like xxx_test.go
//The test function must start with the word Test
//The test function takes one argument only t *testing.T
//In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file.
//*testing.T is your "hook" into the testing framework, so you can do things like t.Fail()

// varName := value <- Variable declaration in Go
