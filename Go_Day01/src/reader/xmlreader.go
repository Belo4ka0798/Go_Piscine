package reader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type XMLReader Recipe

func (r *XMLReader) ReadDB(content []byte) Recipe {
	err := xml.Unmarshal(content, r)
	if err != nil {
		fmt.Println("Invalid XML")
		os.Exit(1)
	}
	return Recipe(*r)
}

func (r *XMLReader) WriteDB(recipe Recipe) {
	fd, err := os.Create("file.json")
	if err != nil {
		fmt.Println("File not Create!")
		os.Exit(1)
	}
	out, err := json.MarshalIndent(recipe, "", "	")
	if err != nil {
		fmt.Println("File not Marshal!")
		os.Exit(1)
	}
	i, err := fmt.Fprint(fd, string(out))
	if err != nil {
		fmt.Printf("Can not write %d!", i)
		os.Exit(1)
	}
}
