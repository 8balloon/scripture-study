package main

import (
	"fmt"

	"github.com/8balloon/scripture-study/bom"
)

func main() {
	s := bom.Obfuscated("yolo!")
	fmt.Println(string(s))

	cases := bom.Case{
		Variants: []bom.Obfuscated{s},
	}
	fmt.Println(string(cases.Variants[0]))
}
