package main

import (
	"flag"
	"fmt"
	"github.com/r3labs/diff/v3"
	"go_day01/reader"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	useFile1 := flag.String("old", "", "parse file old")
	useFile2 := flag.String("new", "", "parse file new")

	flag.Parse()
	var old, new reader.Recipe
	if *useFile1 != "" && *useFile2 != "" {
		old = getRecipe(*useFile1)
		new = getRecipe(*useFile2)
	} else {
		fmt.Println("Use '--old' and '--new' flag to compare")
	}
	compareRecipe(old, new)
}

func getRecipe(filename string) reader.Recipe {
	var recipe reader.Recipe
	if strings.HasSuffix(filename, ".xml") {
		myStruct := new(reader.XMLReader)
		recipe = parseFile(myStruct, filename)
	} else if strings.HasSuffix(filename, ".json") {
		myStruct := new(reader.JSONReader)
		recipe = parseFile(myStruct, filename)
	} else {
		fmt.Fprint(os.Stderr, "error: invalid file extension\n")
		os.Exit(1)
	}
	return recipe
}

func parseFile(readers reader.DBReader, filename string) reader.Recipe {
	var recipe reader.Recipe
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Println("File not exists!")
		os.Exit(1)
	}
	infile, err := io.ReadAll(fd)
	if err != nil && err != io.EOF {
		fmt.Println("!!!EOF!!!")
	}
	recipe = readers.ReadDB(infile)
	return recipe
}

func compareRecipe(old reader.Recipe, new reader.Recipe) {
	diffLog, _ := diff.Diff(old, new)
	for _, change := range diffLog {
		switch change.Type {
		case diff.CREATE:
			checkAdd(change, new)
			continue
		case diff.UPDATE:
			checkUpd(change, new)
			continue
		case diff.DELETE:
			checkDel(change, old)
			continue
		}
	}

}

func checkAdd(change diff.Change, recipe reader.Recipe) {
	last := change.Path[len(change.Path)-1]
	switch last {
	case "Name":
		fmt.Printf("ADDED cake \"%s\"\n", change.To)
	case "ItemName":
		fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", change.To, getCake(change.Path, recipe))
	case "ItemUnit":
		fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
			change.To, getIngredient(change.Path, recipe), getCake(change.Path, recipe))
	}
}

func checkUpd(change diff.Change, recipe reader.Recipe) {
	last := change.Path[len(change.Path)-1]
	switch last {
	case "Time":
		fmt.Printf("CHANGED cooking time for cake \"%s\" – from \"%s\" to \"%s\"\n",
			getCake(change.Path, recipe), change.From, change.To)
	case "ItemCount":
		fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" – from \"%s\" to \"%s\"\n",
			getIngredient(change.Path, recipe), getCake(change.Path, recipe), change.From, change.To)
	case "ItemUnit":
		if change.To != "" {
			fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" – from \"%s\" to \"%s\"\n",
				getIngredient(change.Path, recipe), getCake(change.Path, recipe), change.From, change.To)
		} else {
			fmt.Printf("REMOVED unit for ingredient \"%s\" for cake \"%s\"\n",
				getIngredient(change.Path, recipe), getCake(change.Path, recipe))
		}
	}
}

func checkDel(change diff.Change, recipe reader.Recipe) {
	path := change.Path
	last := path[len(path)-1]
	switch last {
	case "Name":
		fmt.Printf("REMOVED cake \"%s\"\n", change.From)
		fmt.Println("delete")
	case "ItemName":
		fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n",
			getCake(path, recipe), change.From)
	case "ItemUnit":
		fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
			getIngredient(path, recipe), getCake(path, recipe), change.From)
	}
}

func getCake(path []string, recipe reader.Recipe) string {
	cake, err := strconv.Atoi(path[1])
	if err != nil {
		fmt.Errorf("Cake is not valid")
		os.Exit(1)
	}
	return recipe.Cake[cake].Name
}

func getIngredient(path []string, recipe reader.Recipe) string {
	cake, err := strconv.Atoi(path[1])
	if err != nil {
		fmt.Errorf("Undefined cake!")
		os.Exit(1)
	}
	ingredient, err := strconv.Atoi(path[3])
	if err != nil {
		fmt.Errorf("Undefined cake!")
		os.Exit(1)
	}
	return recipe.Cake[cake].Ingredients[ingredient].ItemName
}
