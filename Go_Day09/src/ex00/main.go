package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func main() {
	fmt.Println("Start")
	// создание канала и по лучение из него данных
	res := sleepSort([]int{1, 5, 7, 9, 4, 2, 4, 2, 5, 7, 8, 3, 2, 3, 2, 1, 6, 3})

	for val := range res {
		fmt.Print(val, " ")
	}
	fmt.Println("\nEnd")
}

func sleepSort(arr []int) <-chan int {
	wg := sync.WaitGroup{}
	ch := make(chan int, len(arr))
	defer close(ch)
	sort.Ints(arr)
	fmt.Println(len(arr))
	for _, val := range arr {
		wg.Add(1)
		go func(wg *sync.WaitGroup, ch chan int, val int) {
			defer wg.Done()
			ch <- val
		}(&wg, ch, val)
		time.Sleep(time.Second / 10)

	}
	wg.Wait()
	return ch
}
