package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Input Flags!")
		os.Exit(1)
	}
	l := flag.Bool("l", false, "show count lines only")
	w := flag.Bool("w", false, "show words only")
	s := flag.Bool("m", false, "show symbols only")
	flag.Parse()

	if (*l && (*w || *s)) || (*w && (*l || *s)) || (*s && (*l || *w)) {
		fmt.Println("Many flags! (flags > 1)")
		os.Exit(1)
	}
	switch {
	case !*l && !*w && !*s:
		Start("words", 1)
	case *l:
		Start("lines", 2)
	case *w:
		Start("words", 2)
	case *s:
		Start("symbols", 2)
	}
}

func Start(key string, iter int) {
	var wg sync.WaitGroup
	for _, file := range os.Args[iter:] {
		wg.Add(1)
		go Change(&wg, key, file)
	}
	wg.Wait()
}

func Change(wg *sync.WaitGroup, key string, file string) {
	defer wg.Done()
	fs, err := os.Open(file)
	if err != nil {
		fmt.Println("File not open!")
	}
	scan := bufio.NewScanner(fs)
	switch key {
	case "lines":
		scan.Split(bufio.ScanLines)
	case "words":
		scan.Split(bufio.ScanWords)
	case "symbols":
		scan.Split(bufio.ScanRunes)
	}
	Count(file, fs, scan)
	fs.Close()
}

func Count(path string, fs *os.File, scan *bufio.Scanner) int {
	var res int
	for scan.Scan() {
		res++
	}
	if err := scan.Err(); err != nil {
		fmt.Println("File is not valid!")
	}
	fs.Close()
	fmt.Printf("%d %s\n", res, path)
	return res
}
