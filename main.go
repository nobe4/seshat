package main

import (
	"fmt"
	"time"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
	"github.com/tdewolff/canvas/renderers/pdf"
)

const lorem string = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam in dui mauris. Vivamus hendrerit arcu sed`

func main() {
	fmt.Printf("Running at %s\n", time.Now().Format("2006-01-02 15:04:05"))

	c := canvas.New(210, 297)
	ctx := canvas.NewContext(c)

	karrikFont := canvas.NewFontFamily("Karrik")
	if err := karrikFont.LoadFontFile("./font/OTF/Lineal-Regular.otf", canvas.FontRegular); err != nil {
		panic(err)
	}

	size := 30.0
	for y := 297.0; y > 0; {
		face := karrikFont.Face(size, canvas.Black, canvas.FontRegular, canvas.FontNormal)

		txt := canvas.NewTextBox(face, lorem, 210.0, 0.0, canvas.Left, canvas.Top, 0.0, 0.0)
		ctx.DrawText(0, y, txt)

		y -= txt.Bounds().H
		size -= 1.0
	}
	if err := renderers.Write("test.pdf", c, &pdf.Options{}); err != nil {
		panic(err)
	}
}
