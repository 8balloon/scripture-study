package main

import (
	"math"
)

type Wing struct {
	Size     int
	Before   map[rune]int
	After    map[rune]int
	Symmetry float64
}

type WingSet map[rune]Wing

func CosineSim(A []int, B []int) float64 {
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

func (wset WingSet) Read(str string) {
	runes := []rune(str)
	lastIdx := len(runes) - 1
	for i, r := range runes {
		w, ok := wset[r]
		if !ok {
			var wing Wing
			wing.Before = make(map[rune]int)
			wing.After = make(map[rune]int)
			w = wing
			// need to assign for every wing
			// because otherwise we're just modiying
			// a copy of the map contents
			// wset[r]= w
		}
		w.Size++
		if i > 0 {
			prev := runes[i-1]
			w.Before[prev]++
		}
		if i < lastIdx {
			next := runes[i+1]
			w.After[next]++
		}
		wset[r] = w
	}

	for r, wing := range wset {
		runeSet := make(map[rune]bool)
		for r := range wing.Before {
			runeSet[r] = true
		}
		for r := range wing.After {
			runeSet[r] = true
		}
		var beforeVec []int
		var afterVec []int
		for r := range runeSet {
			beforeVec = append(beforeVec, wing.Before[r])
			afterVec = append(afterVec, wing.After[r])
		}
		wing.Symmetry = CosineSim(beforeVec, afterVec)
		wset[r] = wing
	}
}

func CreateWingSet(s string) WingSet {
	wset := make(WingSet)
	wset.Read(s)
	return wset
}
