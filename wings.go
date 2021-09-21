package main

type Wings struct {
	Before map[rune]int
	After  map[rune]int
}

type WingSet map[rune]Wings

// use io.Reader/Writer?
func GetWingSet(str string) WingSet {
	runes := []rune(str)
	lastIdx := len(runes) - 1
	ws := make(WingSet)
	for i, r := range runes {
		w := ws[r]
		if i > 0 {
			prev := runes[i-1]
			w.Before[prev]++
		}
		if i < lastIdx {
			next := runes[i+1]
			w.After[next]++
		}
	}
	return ws
}
