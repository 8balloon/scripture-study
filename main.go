package main

import (
	"fmt"

	"github.com/8balloon/scripture-study/bom"
)

func main() {
	s := bom.OString("yolo!")
	fmt.Println(string(s))

	cases := bom.LetterType{
		Variants: []bom.ORune{'Y'},
	}
	fmt.Println(string(cases.Variants[0]))
}
