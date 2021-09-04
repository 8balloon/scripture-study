package bom

type ORune rune

const shiftlen = 58

// create an obfuscated rune, or ORune
// we obfuscate runes to reduce preconceptions
// about how text should be processed
func CreateORune(r rune) ORune {
	return ORune(r + shiftlen)
}

// deobfuscates an ORune
func (or ORune) Deobfuscate() rune {
	return rune(or) - shiftlen
}
