package main

import "fmt"

func main() {

	fmt.Println(AssumingDelimiter(" ", "hello, world"))

	withPunc := AssumingPunctuation(",", "Hello, world")
	fmt.Println(len(withPunc), withPunc)

	scriptures := GetScriptures()
	pogp := scriptures[len(scriptures)-1]
	fmt.Println(pogp.JSON.Books[0].Book)
	fmt.Println(scriptures[0].JSON.Books[0].Book)
}
