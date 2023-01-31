package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type inMap map[int]int

type flags struct {
	def    int
	mean   int
	median int
	mode   int
	sd     int
}

func main() {
	Start()
}

func Start() error {
	var m inMap = make(map[int]int)
	var f flags
	var arr []int
	count := 0
	err := ValidArgc(&f)
	if err != nil {
		fmt.Println("Invalid Arguments!!")
		return nil
	}
	fmt.Println("Enter \"stop\" to finish entering the sequence!")
	for {
		var input string
		fmt.Scan(&input)
		if input == "stop" {
			break
			fmt.Println("OK!")
		}
		num, err := strconv.Atoi(input)
		if err != nil || num > 100000 || num < -100000 {
			fmt.Println("Invalid input!")
			return err
		}
		arr = append(arr, num)
		m[num] += 1
		count++
	}
	if len(arr) == 0 {
		fmt.Println("Sequence is empty!")
		return nil
	}
	Show(arr, m, f, count)
	return nil
}

func Mean(arr []int, count int) float64 {
	var mean float64 = 0.0
	for _, value := range arr {
		mean += float64(value)
	}
	mean = mean / float64(count)
	return mean
}

func Median(arr []int) float64 {
	sort.Ints(arr)
	if len(arr)%2 == 1 {
		return float64(arr[len(arr)/2])
	} else {
		return 0.5 * (float64(arr[len(arr)/2.0-1.0]) + float64(arr[len(arr)/2.0]))
	}
}

func Mode(m inMap) int {
	var res = 0
	var beforeVal int = 0
	for key, value := range m {
		if beforeVal < value {
			res = key
			beforeVal = value
		} else if beforeVal == value && res > key {
			res = key
		}
	}
	return res
}

func SD(arr []int, count int) float64 {
	sd := 0.0
	var mean = Mean(arr, count)
	for i := 0; i < count; i++ {
		sd += math.Pow(float64(arr[i])-mean, 2)
	}
	sd /= float64(count)
	return math.Sqrt(sd)
}

func Show(arr []int, m inMap, f flags, count int) {
	if f.def == 1 {
		fmt.Printf("Mean: %.2f\n", Mean(arr, count))
		fmt.Printf("Median: %.2f\n", Median(arr))
		fmt.Printf("Mode: %d\n", Mode(m))
		fmt.Printf("SD: %.2f\n", SD(arr, count))
	} else {
		if f.mean == 1 {
			fmt.Printf("Mean: %.2f\n", Mean(arr, count))
		}
		if f.median == 1 {
			fmt.Printf("Median: %.2f\n", Median(arr))
		}
		if f.mode == 1 {
			fmt.Printf("Mode: %d\n", Mode(m))
		}
		if f.sd == 1 {
			fmt.Printf("SD: %.2f\n", SD(arr, count))
		}
	}
}

func ValidArgc(flags *flags) error {
	args := os.Args
	if len(args) == 1 {
		flags.def = 1
		return nil
	}
	for c, flag := range args {
		if c == 0 {
			continue
		}
		switch flag {
		case "mean":
			flags.mean = 1
			break
		case "median":
			flags.median = 1
			break
		case "mode":
			flags.mode = 1
			break
		case "sd":
			flags.sd = 1
			break
		default:
			return fmt.Errorf("undefind flag")
			break
		}
	}
	return nil
}
