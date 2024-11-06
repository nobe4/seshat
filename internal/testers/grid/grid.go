package grid

import (
	"math"

	"github.com/nobe4/seshat/internal/config"
	"github.com/nobe4/seshat/internal/font"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
)

type Box struct {
	text *canvas.Text
	w    float64
	h    float64
}

func Test(pdf *pdf.PDF, fonts font.Fonts, config config.Config, rule config.Render) {
	width, height := pdf.Size()

	columns := rule.Rules.Columns

	gridSize := biggestGridSize(len(fonts))
	boxes := []Box{}
	maxW, maxH := 0.0, 0.0

	// TODO: do a binary search
	// Find the smallest font size that fits the text in the grid.
	size := rule.Rules.Size

	for {
		boxes, maxW, maxH = prepareBoxes(size, fonts, rule.Rules.Features, rule.Inputs)

		if float64(columns)*maxW*float64(gridSize) > width-10 {
			size -= 1
			continue
		}

		break
	}

	y := height - maxH
	currentColumn := -1

	fontH := maxH * float64(gridSize)
	can := canvas.New(width, height)
	ctx := canvas.NewContext(can)

	for i := range len(rule.Inputs) {
		currentColumn += 1

		// Column break
		if currentColumn >= columns {
			currentColumn = 0
			y -= maxH * float64(gridSize)
		}

		nextY := y - fontH

		// page break
		if nextY < 0 {
			can.RenderTo(pdf)
			can = canvas.New(width, height)
			ctx = canvas.NewContext(can)
			pdf.NewPage(width, height)

			y = height - maxH
		}

		x := float64(currentColumn) * maxW * float64(gridSize)

		for j := range len(fonts) {
			b := boxes[j+i*len(fonts)]

			ctx.DrawText(
				x+float64(j%gridSize)*maxW,
				y-float64(j/gridSize)*maxH,
				b.text,
			)
		}

	}

	can.RenderTo(pdf)
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
