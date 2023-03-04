package main

import (
	"container/heap"
	"errors"
	"fmt"
	"log"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap []Present

func main() {
	var n int = 5
	presents := []Present{
		Present{Value: 5, Size: 2},
		Present{Value: 4, Size: 5},
		Present{Value: 3, Size: 1},
		Present{Value: 5, Size: 3},
		Present{Value: 5, Size: 1},
		Present{Value: 7, Size: 4},
	}
	colPre, err := getNCoolestPresents(presents, n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(colPre)
}

func getNCoolestPresents(pres []Present, n int) ([]Present, error) {
	h := PresentHeap(pres)
	if n > len(pres) {
		return nil, errors.New("`n` is larger than the size of the slice or is negative")
	}
	heap.Init(&h)
	return h[:n], nil
}

func (ph *PresentHeap) Pop() any {
	old := *ph
	n := len(old)
	x := old[n-1]
	*ph = old[:n-1]
	return x
}

func (ph *PresentHeap) Push(x any) {
	*ph = append(*ph, x.(Present))
}

func (ph PresentHeap) Len() int {
	return len(ph)
}

func (ph PresentHeap) Less(i, j int) bool {
	if ph[i].Value == ph[j].Value {
		return ph[i].Size < ph[j].Size
	}
	return ph[i].Value > ph[j].Value
}

func (ph PresentHeap) Swap(i, j int) {
	ph[i], ph[j] = ph[j], ph[i]
}
