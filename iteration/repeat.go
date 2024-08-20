package iteration

import "fmt"

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func Main() {
	fmt.Println("Hey there!")
}
