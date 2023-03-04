package main

import (
	"fmt"
	"log"
	"math"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap []Present

func main() {
	var cap int = 3
	presents := []Present{
		Present{Value: 5, Size: 2},
		Present{Value: 4, Size: 5},
		Present{Value: 3, Size: 1},
		Present{Value: 5, Size: 3},
		Present{Value: 5, Size: 1},
		Present{Value: 7, Size: 4},
	}
	//var cap int = 15
	//var presents = []Present{{1, 4}, {7, 0}, {0, 3}, {0, 0}, {18, 18}, {1, 3}, {1, 3}, {1, 3}, {44, 1}, {111, 7}, {12, 12}, {5, 6}}
	for i := range presents {
		if presents[i].Size < 1 {
			log.Fatal("Size can not < 1")
		}
	}
	fmt.Printf("Alex = %v\n", grabPresents(presents, cap))

}

func grabPresents(pres []Present, cap int) (res []Present) {
	l := len(pres)
	arr := make([][]int, l+1)
	for i := range arr {
		arr[i] = make([]int, cap+1)
	}
	for i := 1; i <= l; i++ {
		for j := 0; j <= cap; j++ {
			if i == 0 || j == 0 {
				arr[i][j] = 0
			} else {
				if pres[i-1].Size > j {
					arr[i][j] = arr[i-1][j]
				} else {
					prev := arr[i-1][j]
					formula := pres[i-1].Value + arr[i-1][j-pres[i-1].Size]
					arr[i][j] = int(math.Max(float64(prev), float64(formula)))
				}
			}
		}
	}
	i, j := l, cap
	for i > 0 && j > 0 {
		if arr[i][j] == arr[i-1][j] {
			i--
		} else {
			res = append(res, pres[i-1])
			j -= pres[i-1].Size
			i--
		}
	}
	return res
}
