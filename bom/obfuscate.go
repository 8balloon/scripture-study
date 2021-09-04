package bom

type Obfuscated string

const shiftlen = 58

// obfuscates a string to reduce preconceptions
// about how text should be processed
func Obfuscate(s string) Obfuscated {
	runes := []rune(s)
	for i, rune := range runes {
		runes[i] = rune + shiftlen
	}
	return Obfuscated(runes)
}

// deobfuscates a string
func Deobfuscate(o Obfuscated) string {
	runes := []rune(string(o))
	for i, rune := range runes {
		runes[i] = rune - shiftlen
	}
	return string(runes)
}
