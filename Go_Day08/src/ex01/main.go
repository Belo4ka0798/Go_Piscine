package main

func main() {

}

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func describePlant(p any) {
	Type().Field()
}
