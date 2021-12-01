// Writing a program in Go, you will have a main package defined with a main func inside it.
// Packages are ways of grouping up related Go code together
package main

// importing a package which contains the Println function that we use to print
import "fmt"

// // func keyword is how we define a function with a name and a body
// func main() {
//     fmt.Println("Hello, world")
// }

// new function but added a keyword string in the definition. This means the function returns a string.
func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}