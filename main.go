package main

import "fmt"

func main() {
	pogpJSON := GetPOGPSJ()

	fmt.Println(pogpJSON.Books[0].Book)

	fmt.Println(AssumingDelimiter(" ", "hello, world"))

	withPunc := AssumingPunctuation(",", "Hello, world")
	fmt.Println(len(withPunc), withPunc)
}
