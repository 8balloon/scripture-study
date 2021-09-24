package main

import "testing"

func TestCosineSim(t *testing.T) {
	r1 := CosineSim([]int{1, 1, 0}, []int{2, 0, 2})
	t.Log(r1)
	r2 := CosineSim([]int{1, 1, 1}, []int{2, 2, 1})
	t.Log(r2)
}

func TestWingSet(t *testing.T) {
	ws := CreateWingSet("Helllo")
	t.Logf("%v", ws)
	t.Logf("%v", ws[32])
}
