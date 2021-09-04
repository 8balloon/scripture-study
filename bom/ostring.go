package bom

type OString string

// create an obfuscated string, or OString
// we obfuscate strings to reduce preconceptions
// about how text should be processed
func CreateOString(s string) OString {
	runes := []rune(s)
	for i, r := range runes {
		runes[i] = rune(CreateORune(r))
	}
	return OString(string(runes))
}

// deobfuscates an OString
func (os OString) Deobfuscate() string {
	runes := []rune(string(os))
	for i, r := range runes {
		runes[i] = rune(ORune(r).Deobfuscate())
	}
	return string(runes)
}
