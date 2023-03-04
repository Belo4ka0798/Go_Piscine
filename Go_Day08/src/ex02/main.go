package main

import (
	g "github.com/AllenDang/giu"
)

func loop() {

	g.SingleWindow().Layout(
		g.Label("School 21"),
	)
}

func main() {
	wnd := g.NewMasterWindow("School 21", 300, 200, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
