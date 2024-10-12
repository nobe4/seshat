package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
)

const lorem string = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam in dui mauris. Vivamus hendrerit arcu sed`

type FontMap struct {
	name  string
	style canvas.FontStyle
}

func main() {
	fmt.Printf("Running at %s\n", time.Now().Format("2006-01-02 15:04:05"))

	f, err := os.Create("test.pdf")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	font := canvas.NewFontFamily("Lineal")
	pdf := pdf.New(f, 210, 297, &pdf.Options{})

	fontThicknesses := []FontMap{
		{"Thin", canvas.FontThin},
		{"Light", canvas.FontLight},
		{"Regular", canvas.FontRegular},
		{"Medium", canvas.FontMedium},
		{"Bold", canvas.FontBold},
		{"Black", canvas.FontBlack},
		{"Heavy", canvas.FontExtraBold},
	}

	y := 297.0
	for _, s := range fontThicknesses {
		if err := font.LoadFontFile("./font/OTF/Lineal-"+s.name+".otf", s.style); err != nil {
			panic(err)
		}

		c := canvas.New(210, 297)
		ctx := canvas.NewContext(c)

		size := 30.0
		// y := 297.0
		// for y, i := 297.0, 20; y > 0 && i > 0; i-- {
		face := font.Face(size, canvas.Black, s.style, canvas.FontNormal)

		txt := canvas.NewTextBox(face, s.name+lorem, 210.0, 0.0, canvas.Left, canvas.Top, 0.0, 0.0)
		ctx.DrawText(0, y, txt)

		y -= txt.Bounds().H
		size -= 1.5
		// }

		c.RenderTo(pdf)
		// pdf.NewPage(210, 297)
	}

	if err := pdf.Close(); err != nil {
		panic(err)
	}
}
