// go run hello_world.go => Run the script
// go build hello_world.go => build the binary of the script and then ./hello-world to execute that binary

package main

import "fmt"

func hello() string {
	return "hello world"
}

func main() {
	fmt.Println(hello())
}

// separation of side effects from the domain code
