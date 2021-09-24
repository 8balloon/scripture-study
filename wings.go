package main

import "math"

type Wings struct {
	Size     int
	Before   map[rune]int
	After    map[rune]int
	Symmetry float64
}

type WingSet map[rune]Wings

func Cosinesim(A []int, B []int) float64 {
	dotproduct := 0
	var mA float64
	var mB float64
	for i, va := range A {
		vb := B[i]
		dotproduct += (va * vb)
		mA += float64(va * va)
		mB += float64(vb * vb)
	}
	mA = math.Sqrt(mA)
	mB = math.Sqrt(mB)
	similarity := float64(dotproduct) / ((mA) * (mB))
	return similarity
}

// use io.Reader/Writer?
func GetWingSet(str string) WingSet {
	runes := []rune(str)
	lastIdx := len(runes) - 1
	ws := make(WingSet)
	for i, r := range runes {
		w := ws[r]
		w.Size++
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
