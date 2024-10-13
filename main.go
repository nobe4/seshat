package main

import (
	"fmt"
	"os"
	"test/internal/font"
	"time"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
)

const (
	fontDir = "./font/OTF"
	lorem   = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam in dui mauris. Vivamus hendrerit arcu sed`
	outFile = "test.pdf"

	// pdf size in points
	width  = 210.0
	height = 297.0
)

func main() {
	start := time.Now()
	fmt.Printf("Start at %s\n", start.Format("15:04:05"))

	f, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pdf := pdf.New(f, width, height, &pdf.Options{})

	fonts, err := font.LoadFromDir(fontDir)
	if err != nil {
		panic(err)
	}

	loopSizeAndWeightWithFontName(pdf, fonts)
	loopSizeAndWeightWithLorem(pdf, fonts)

	if err := pdf.Close(); err != nil {
		panic(err)
	}

	end := time.Now()
	fmt.Printf("Ran at %s in %fs\n", end.Format("15:04:05"), end.Sub(start).Seconds())
}

func loopSizeAndWeightWithFontName(pdf *pdf.PDF, fonts font.Fonts) {
	for _, font := range fonts {
		c := canvas.New(width, height)
		ctx := canvas.NewContext(c)

		size := 50.0
		for y, i := height, 30; y > 0 && i > 0; i-- {
			face := font.Font.Face(size, canvas.Black)

			txt := canvas.NewTextBox(face, font.Name, width, 0.0, canvas.Left, canvas.Top, 0.0, 0.0)
			ctx.DrawText(0, y, txt)

			y -= txt.Bounds().H
			size -= 1.6
		}

		c.RenderTo(pdf)

		pdf.NewPage(width, height)
	}
}

func loopSizeAndWeightWithLorem(pdf *pdf.PDF, fonts font.Fonts) {
	for _, font := range fonts {
		c := canvas.New(width, height)
		ctx := canvas.NewContext(c)

		size := 30.0
		for y, i := height, 30; y > 0 && i > 0; i-- {
			face := font.Font.Face(size, canvas.Black)

			txt := canvas.NewTextBox(face, lorem, width, 0.0, canvas.Left, canvas.Top, 0.0, 0.0)
			ctx.DrawText(0, y, txt)

			y -= txt.Bounds().H
			size -= 1.6
		}

		c.RenderTo(pdf)

		pdf.NewPage(width, height)
	}
}
