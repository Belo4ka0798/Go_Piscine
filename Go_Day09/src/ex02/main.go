package main

import (
	"fmt"
	"sync"
)

func fillChanel(a []int) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		for _, val := range a {
			out <- val
		}
		close(out)

	}()
	return out
}

func main() {
	ch1 := fillChanel([]int{1, 2, 3, 4, 5, 6})
	ch2 := fillChanel([]int{1, 2, 3, 4, 5, 6})

	for i := range multiplex(ch1, ch2) {
		fmt.Println(i)
	}
}

func multiplex(channels ...<-chan interface{}) <-chan interface{} {
	result := make(chan interface{})
	wg := sync.WaitGroup{}

	import1 := func(item <-chan interface{}) {
		for i := range item {
			result <- i
		}
		wg.Done()
	}

	wg.Add(len(channels))

	for _, res := range channels {
		go import1(res)
	}

	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}
