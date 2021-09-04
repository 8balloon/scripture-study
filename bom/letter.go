package bom

// a Letter is an abstract rune which can exist as one or more forms
// (e.g., 'a' + 'A')
type LetterType struct {
	Variants []ORune
}

type Letter struct {
	LetterType *LetterType
	ORune      ORune
}
