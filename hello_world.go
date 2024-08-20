// go run hello_world.go => Run the script
// go build hello_world.go => build the binary of the script and then ./hello-world to execute that binary

package main

import "fmt"

func hello() string {
	return "hello world"
}

func hello2(name string) string {
	if name != "" {
		return "Hello, " + name
	}
	return "Hello, Stranger!"
}

func main() {
	fmt.Println(hello())
	fmt.Println(hello2(""))
}

// separation of side effects from the domain code
