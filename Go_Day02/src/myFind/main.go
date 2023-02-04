package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Input Flags!")
		os.Exit(1)
	}
	files := flag.Bool("f", false, "show files only")
	dirs := flag.Bool("d", false, "show directories only")
	links := flag.Bool("sl", false, "show symbol links only")
	ext := flag.String("ext", "", "files extension")
	flag.Parse()

	if *ext != "" && !*files {
		fmt.Println("\"-ext\" not use without \"-f\"")
		os.Exit(1)
	}
	if !*files && !*dirs && !*links {
		MyFind(os.Args[1], true, true, true, *ext)
	} else {
		MyFind(os.Args[len(os.Args)-1], *files, *dirs, *links, *ext)
	}
}

func MyFind(path string, f, d, sl bool, ext string) {
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	dirs, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("%v", err)
	}
	for _, s := range dirs {
		if s.IsDir() {
			if d {
				fmt.Println(path + s.Name())
			}
			buff := path + s.Name()
			MyFind(buff, f, d, sl, ext)
		} else {
			link, err := os.Readlink(path + s.Name())
			if err == nil {
				_, errLink := os.Open(path + link)
				if sl && errLink != nil {
					fmt.Printf("%s -> %s\n", path+s.Name(), "[broken]")
				} else if sl {
					fmt.Printf("%s -> %s\n", path+s.Name(), link)
				}
			} else if f && strings.HasSuffix(s.Name(), ext) {
				fmt.Println(path + s.Name())
			}
		}
	}
}
