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
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	out, err := json.MarshalIndent(recipe, "", "	")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprint(fd, string(out))
}
