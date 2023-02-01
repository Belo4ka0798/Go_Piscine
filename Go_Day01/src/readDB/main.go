package main

import (
	"flag"
	"fmt"
	"go_day01/reader"
	"io"
	"os"
	"strings"
)

func main() {
	useFile := flag.String("f", "", "parse file")
	flag.Parse()

	if *useFile != "" {
		Start(*useFile)
	} else {
		fmt.Println("Use '-f' flag to pass argument")
	}
}

func Start(filename string) {
	if strings.HasSuffix(filename, ".xml") {
		myStruct := new(reader.XMLReader)
		parseFile(myStruct, filename)
	} else if strings.HasSuffix(filename, ".json") {
		myStruct := new(reader.JSONReader)
		parseFile(myStruct, filename)
	} else {
		fmt.Println("error: invalid file extension")
		os.Exit(1)
	}
}

func parseFile(readers reader.DBReader, filename string) {
	var recipe reader.Recipe
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Println("File not exists!")
		os.Exit(1)
	}
	infile, err := io.ReadAll(fd)
	if err != nil && err != io.EOF {
		fmt.Println("")
	}
	recipe = readers.ReadDB(infile)
	readers.WriteDB(recipe)
}
