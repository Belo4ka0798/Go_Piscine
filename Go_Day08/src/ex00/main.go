package main

import (
	"fmt"
	"log"
	"unsafe"
)

func main() {
	res, _ := getElement([]int{1, 2, 3, 4, 5, 6}, 5)
	fmt.Println(res)
}

func getElement(arr []int, idx int) (int, error) {
	if idx < 0 {
		log.Fatal("Err: Index < 0!")
	}
	if idx > len(arr) {
		log.Fatal("Err: Index out of range!")
	}
	size := unsafe.Sizeof(int(0))
	p := uintptr(unsafe.Pointer(&arr[0]))

	return (*(*int)(unsafe.Pointer(size*uintptr(idx) + p))), nil
}
