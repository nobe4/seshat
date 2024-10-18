package grid

import (
	"fmt"
	"math"

	"github.com/nobe4/seshat/internal/font"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
)

type Box struct {
	text *canvas.Text
	w    float64
	h    float64
	x    float64
	y    float64
}

func Test(pdf *pdf.PDF, fonts font.Fonts, features string, inputs []string) {
	width, height := pdf.Size()

	gridSize := biggestGridSize(len(fonts))
	boxes := []Box{}
	maxW, maxH := 0.0, 0.0

	// TODO: do a binary search
	// Find the smallest font size that fits the text in the grid.
	size := fonts[0].Size
	for {
		boxes, maxW, maxH = prepareBoxes(size, fonts, features, inputs)

		if maxW*float64(gridSize) > width-10 {
			size -= 1
			continue
		}

		break
	}

	y := height - maxH
	fontH := maxH * float64(gridSize)
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	i := 0
	for range len(inputs) {
		nextY := y - fontH
		fmt.Printf("fontH: %f  nextY: %f, y: %f\n", fontH, nextY, y)

		// page break
		if nextY < 0 {
			c.RenderTo(pdf)
			c = canvas.New(width, height)
			ctx = canvas.NewContext(c)
			pdf.NewPage(width, height)

			y = height - maxH
		}

		for range len(fonts) {
			b := boxes[i]
			fmt.Printf("i: %d, x: %d, y: %f\n", i, i%gridSize, i/gridSize)

			b.x = float64(i%gridSize) * maxW
			b.y = y

			fmt.Printf("x: %f, y: %f, i: %s\n", b.x, b.y, b.text)
			ctx.DrawText(b.x, b.y, b.text)

			i += 1
			if i%gridSize == 0 {
				y -= maxH
			}
		}

		fmt.Printf("y: %f\n", y)
	}

	c.RenderTo(pdf)

	pdf.NewPage(width, height)
}

func prepareBoxes(size float64, fonts font.Fonts, features string, inputs []string) ([]Box, float64, float64) {
	boxes := []Box{}

	biggestW := 0.0
	biggestH := 0.0

	for _, input := range inputs {
		for _, font := range fonts {
			face := font.Font.Face(size, canvas.Black)
			face.Font.SetFeatures(features)

			txt := canvas.NewTextLine(face, input, canvas.Left)

			boxes = append(boxes, Box{
				text: txt,
				w:    txt.Width,
				h:    txt.Height,
			})
			if txt.Width > biggestW {
				biggestW = txt.Width
			}
			if txt.Height > biggestH {
				biggestH = txt.Height
			}
		}
	}

	return boxes, biggestW, biggestH
}

func biggestGridSize(l int) int {
	return int(math.Ceil(math.Sqrt(float64(l))))
}
