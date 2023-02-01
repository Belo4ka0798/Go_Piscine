package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	useFile1 := flag.String("old", "", "parse file old")
	useFile2 := flag.String("new", "", "parse file new")

	flag.Parse()
	if *useFile1 != "" && *useFile2 != "" {
		old := openFile(*useFile1)
		nev := openFile(*useFile2)
		for _, v := range compareFile(nev, old) {
			fmt.Println("ADDED", v)
		}
		for _, v := range compareFile(old, nev) {
			fmt.Println("REMOVED", v)
		}
	} else {
		fmt.Println("Use '--old' and '--new' flag to compare")
	}
}

func openFile(filename string) []string {
	if !strings.HasSuffix(filename, ".txt") {
		fmt.Println("Format is not \"txt\"")
		os.Exit(1)
	}
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error open file!")
		os.Exit(1)
	}
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {
			fmt.Println("File not close!")
		}
	}(fd)
	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	return txtlines
}

func compareFile(old, new []string) []string {
	mb := make(map[string]struct{}, len(new))
	for _, x := range new {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range old {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
