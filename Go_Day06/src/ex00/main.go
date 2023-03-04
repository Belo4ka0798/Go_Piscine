package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	const height, width = 300, 300

	logo := image.NewNRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			logo.Set(x, y, color.NRGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 0,
			})
			radius := (width/2-x)*(width/2-x) + (height/2-y)*(height/2-y)
			if radius < 9000 {
				logo.Set(x, y, color.NRGBA{
					R: 200,
					G: uint8(255 - 255*radius/9000),
					B: 255,
					A: 180,
				})
			}
		}
	}
	f, err := os.Create("logo.png")
	if err != nil {
		return
	}
	defer f.Close()
	err = png.Encode(f, logo)
	if err != nil {
		return
	}
}
