// To run test you need to type in go test in the terminal.
// However before you can go that you would need to have a module file.
// you will need to run go mod init SOMENAME in each new folder before running commands like go test or go build

// Writing tests in Go

// The test needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T
// In order to use the *testing.T type you need to import "testing"
// the t of type *testing.T is your "hook" in the testing framework, so you can do things like t.Fail() when you want to fail.

package main

import "testing"

// func TestHello(t *testing.T) {
// 	got := Hello()
// 	want := "Hello, world"

// 		if got != want {
// 			// calling the Errorf method on our t which will print out a message and fail the test. 
// 			// The f stands for format which allows us to build a string with values inserted into the placeholder values "%q."
// 			t.Errorf("got %q want %q", got, want)
// 		}
// }

// func TestHello(t *testing.T) {
// 	got := Hello("Chris")
// 	want := "Hello, Chris"
// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// }

// refactored our assertion from the original code into a function. Reduces duplication and improves readability of our tests.
// In Go, you can delcare functions inside other functions and assign them to variables. You can call them like normal functions.
// For Helper functions, good idea to use testing.TB which is an interface that *testing.T and *testing B both satisfy, so you can call helper functions from a test or a benchmark.
// t.Helper() is needed to test the test suite that this method is a helper. When a test fails the line number reported will be in our function call rather than inside our test helper.
func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T){
		got := Hello("Jules", "French")
		want := "Bonjour, Jules"
		assertCorrectMessage(t, got, want)
	})
}


