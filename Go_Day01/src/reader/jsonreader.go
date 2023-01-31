package reader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type JSONReader Recipe

func (r *JSONReader) ReadDB(content []byte) Recipe {
	err := json.Unmarshal(content, r)
	if err != nil {
		fmt.Println("Invalid JSON")
		os.Exit(1)
	}
	return Recipe(*r)
}

func (r *JSONReader) WriteDB(recipe Recipe) {
	fd, err := os.Create("file.xml")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	out, err := xml.MarshalIndent(recipe, "", "	")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprint(fd, string(out))
}
