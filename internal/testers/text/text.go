package text

import (
	"github.com/nobe4/seshat/internal/config"
	"github.com/nobe4/seshat/internal/font"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
)

func Test(pdf *pdf.PDF, fonts font.Fonts, config config.Config, rule config.Rule) {
	width, height := pdf.Size()

	y := height
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	for _, input := range rule.Inputs {
		for _, font := range fonts {
			face := font.Font.Face(rule.Args.Size, canvas.Black)
			face.Font.SetFeatures(rule.Args.Features)

			txt := canvas.NewTextBox(face, input, width, 0.0, canvas.Left, canvas.Top, 0.0, 0.0)

			nextY := y - txt.Bounds().H

			if nextY < 0 {
				c.RenderTo(pdf)
				c = canvas.New(width, height)
				ctx = canvas.NewContext(c)
				pdf.NewPage(width, height)

				y = height
			}

			ctx.DrawText(0, y, txt)
			y -= txt.Bounds().H
		}
	}

	c.RenderTo(pdf)

	pdf.NewPage(width, height)
}
