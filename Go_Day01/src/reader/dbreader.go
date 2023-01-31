package reader

type DBReader interface {
	ReadDB(content []byte) Recipe
	WriteDB(recipe Recipe)
}

type Item struct {
	ItemName  string `xml:"itemname" json:"ingredient_name"`
	ItemCount string `xml:"itemcount" json:"ingredient_count"`
	ItemUnit  string `xml:"itemunit" json:"ingredient_unit,omitempty"`
}

type Cake struct {
	Name        string `xml:"name" json:"name"`
	Time        string `xml:"stovetime" json:"time"`
	Ingredients []Item `xml:"ingredients>item" json:"ingredients"`
}

type Recipe struct {
	Cake []Cake `xml:"cake" json:"cake"`
}
