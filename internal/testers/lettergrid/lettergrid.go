package lettergrid

import (
	"fmt"
	"math"
	"test/internal/font"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
)

type Letter struct {
	text *canvas.Text
	w    float64
	h    float64
	x    float64
	y    float64
}

func Test(pdf *pdf.PDF, fonts font.Fonts) {
	for _, l := range "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" {
		letter(pdf, fonts, string(l))
	}
	letter(pdf, fonts, "test")
}

func letter(pdf *pdf.PDF, fonts font.Fonts, letter string) {
	letters := make([]Letter, 0)

	size := 100.0
	gridSize := biggestGridSize(len(fonts))

	for i, font := range fonts {
		face := font.Font.Face(size, canvas.Black)
		txt := canvas.NewTextLine(face, letter, canvas.Left)

		fmt.Printf("x %f, y %f\n", float64(i%gridSize), float64(i/gridSize))
		letters = append(letters, Letter{
			text: txt,
			w:    txt.Width,
			h:    txt.Height,
			x:    float64(i % gridSize),
			y:    float64(i / gridSize),
		})
	}

	biggestW := 0.0
	biggestH := 0.0
	for _, l := range letters {
		if l.w > biggestW {
			biggestW = l.w
		}
		if l.h > biggestH {
			biggestH = l.h
		}
	}

	width := biggestW * (float64(gridSize) + 0.5)
	height := biggestH * (0.5 + float64(gridSize))
	pdf.NewPage(width, height)

	fmt.Printf("biggestW: %f, biggestH: %f\n", biggestW, biggestH)

	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	for _, l := range letters {
		x := biggestW*0.25 + l.x*biggestW + (biggestW-l.w)/2.0
		y := -biggestH*0.125 + height - (l.y+1.0)*biggestH
		fmt.Printf("x: %f, y: %f\n", x, y)
		ctx.DrawText(x, y, l.text)
	}

	c.RenderTo(pdf)
}

func biggestGridSize(l int) int {
	return int(math.Ceil(math.Sqrt(float64(l))))
}
