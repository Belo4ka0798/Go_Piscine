package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	a := flag.Bool("a", false, "use for \"multi\" archive")
	if len(os.Args) < 2 || (len(os.Args) == 2 && *a) {
		fmt.Println("Invalid arguments!")
		os.Exit(1)
	}
	flag.Parse()

	switch {
	case !*a:
		Start("one", 1)
	case *a:
		Start("multi", 3)
	}
}

func Start(key string, start int) {
	var wg sync.WaitGroup
	for _, file := range os.Args[start:] {
		fmt.Println(file + " " + key)

		if strings.HasSuffix(file, ".log") {
			wg.Add(1)
			//go Change(&wg, key, file)
			go func(wg *sync.WaitGroup, key string, file string) {
				defer wg.Done()
				MultiTar(file, key)
			}(&wg, key, file)

			wg.Wait()
		} else {
			fmt.Println("File without suffix \".log\"!")
			os.Exit(1)
		}
	}
}

func MultiTar(fileName string, key string) {
	files := []string{fileName}

	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Printf("Invalid! %s", fileName)
		os.Exit(1)
	}
	timeStamp := strconv.FormatInt(info.ModTime().Unix(), 10)
	tarName := strings.TrimSuffix(fileName, ".log") + "_" + timeStamp + ".tar.gz"
	if key == "one" {
		newTar, err := os.Create(tarName)
		if err != nil {
			fmt.Println("Can`t create tar!")
			os.Exit(1)
		}
		CreateTar(files, newTar)
	} else if key == "multi" {
		err = os.MkdirAll(strings.TrimPrefix(os.Args[2], "/"), os.ModePerm)
		if err != nil {
			fmt.Println("Can`t create dir!")
			os.Exit(1)
		}
		var tarNamePath string
		if strings.HasSuffix(os.Args[2], "/") {
			tarNamePath = os.Args[2]
		} else {
			tarNamePath = os.Args[2] + "/"
		}
		fmt.Println(tarNamePath + tarName)
		newTar, err := os.Create(strings.TrimPrefix(tarNamePath, "/") + info.Name())
		if err != nil {
			log.Fatalln(err)
		}
		defer newTar.Close()
		CreateTar(files, newTar)
	}
}

func CreateTar(files []string, newTar *os.File) {
	gw := gzip.NewWriter(newTar)
	gw.Name = newTar.Name()
	defer gw.Close()
	tr := tar.NewWriter(gw)
	defer tr.Close()

	for _, file := range files {
		err := addToArchive(tr, file)
		if err != nil {
			fmt.Println("Not Create Archive!")
			os.Exit(1)
		}
	}
}

func addToArchive(tr *tar.Writer, filename string) error {
	fs, err := os.Open(filename)
	if err != nil {
		return err
	}
	info, err := fs.Stat()
	if err != nil {
		return err
	}
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}
	//header.Name = filename
	err = tr.WriteHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(tr, fs)
	if err != nil {
		return err
	}
	return nil
}
