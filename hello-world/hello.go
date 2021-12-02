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
// func Hello() string {
// 	return "Hello, world"
// }

// func main() {
// 	fmt.Println(Hello())
// }

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// the function accepts any argument of type string
func Hello(name string, language string) string {
	if name == "" {
			name = "World"
	}
	
	return greetingPrefix(language) + name
}
// named return value (prefix string). This will create a varaible called prefix in your function.
// It will be assigned the "zero" value. This depends on the type, e.g. int are 0 and for string it is ""
// You can return whatever value is set by calling return rather than return prefix.
// the function name starts with a lowercase letter. In Go public functions start with a captial letter and private functions start with a lowecase.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}