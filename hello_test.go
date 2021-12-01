// To run test you need to type in go test in the terminal.
// However before you can go that you would need to have a module file.
// you will need to run go mod init SOMENAME in each new folder before running commands like go test or go build
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
}

// Writing tests in Go

// The test needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T
// In order to use the *testing.T type you need to import "testing"
// the t of type *testing.T is your "hook" in the testing framework, so you can do things like t.Fail() when you want to fail.
