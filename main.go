package main

import "fmt"

func main() {
	pogpJSON := GetPOGPSJ()

	fmt.Println(pogpJSON.Books[0].Book)
}
