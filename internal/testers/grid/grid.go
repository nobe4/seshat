package grid

import (
	"math"
	"strings"

	"github.com/nobe4/seshat/internal/font"
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

func Test(pdf *pdf.PDF, fonts font.Fonts, letters []string) {
	if len(letters) == 0 {
		letters = strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	}

	for _, l := range letters {
		letter(pdf, fonts, string(l))
	}
}

func letter(pdf *pdf.PDF, fonts font.Fonts, letter string) {
	letters := make([]Letter, 0)

	size := 100.0
	gridSize := biggestGridSize(len(fonts))

	for i, font := range fonts {
		face := font.Font.Face(size, canvas.Black)
		txt := canvas.NewTextLine(face, letter, canvas.Left)

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

	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	for _, l := range letters {
		x := biggestW*0.25 + l.x*biggestW + (biggestW-l.w)/2.0
		y := -biggestH*0.125 + height - (l.y+1.0)*biggestH
		ctx.DrawText(x, y, l.text)
	}

	c.RenderTo(pdf)
}

func biggestGridSize(l int) int {
	return int(math.Ceil(math.Sqrt(float64(l))))
}
