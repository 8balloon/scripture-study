package bom

// a Case is an abstract rune which can exist in one or more forms
// (e.g., 'a' + 'A')
type Case struct {
	Variants []Obfuscated
}
