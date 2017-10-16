package alg

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	var a = []int{9,4,1,5,2,8,7,3,5}
	HoareQuickSort(a)
}

func TestInsertSort(t *testing.T) {
	var a = []int{9,4,1,5,2,8,7,3,5}
	InsertionSort(a)
	for i, v := range a {
		t.Log(i, v)
	}
}