package text

import (
	"github.com/nobe4/seshat/internal/font"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
)

func Test(pdf *pdf.PDF, fonts font.Fonts, features string, inputs []string) {
	width, height := pdf.Size()
	pdf.NewPage(width, height)

	y := height
	for _, input := range inputs {
		c := canvas.New(width, height)
		for _, font := range fonts {

			ctx := canvas.NewContext(c)

			size := 30.0
			face := font.Font.Face(size, canvas.Black)
			face.Font.SetFeatures(features)

			txt := canvas.NewTextBox(face, input, width, 0.0, canvas.Left, canvas.Top, 0.0, 0.0)
			ctx.DrawText(0, y, txt)

			y -= txt.Bounds().H
		}

		c.RenderTo(pdf)
	}
}
