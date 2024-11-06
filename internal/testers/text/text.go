package text

import (
	"github.com/nobe4/seshat/internal/config"
	"github.com/nobe4/seshat/internal/font"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
)

func Test(pdf *pdf.PDF, fonts font.Fonts, config config.Config, rule config.Render) {
	if rule.Rules.Responsive {
		testResponsive(pdf, fonts, config, rule)
		return
	}

	width, height := pdf.Size()

	y := height
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	for _, input := range rule.Inputs {
		for _, font := range fonts {
			face := font.Font.Face(rule.Rules.Size, canvas.Black)
			face.Font.SetFeatures(rule.Rules.Features)

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

func testResponsive(pdf *pdf.PDF, fonts font.Fonts, config config.Config, rule config.Render) {
	width, height := pdf.Size()

	y := height
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	for _, input := range rule.Inputs {
		for _, font := range fonts {

			size := findFontSize(font, input, rule)

			face := font.Font.Face(size, canvas.Black)
			face.Font.SetFeatures(rule.Rules.Features)

			text := canvas.NewTextLine(face, input, canvas.Left)

			nextY := y - text.Bounds().H

			if nextY < 0 {
				c.RenderTo(pdf)
				c = canvas.New(width, height)
				ctx = canvas.NewContext(c)
				pdf.NewPage(width, height)

				y = height
			}

			y -= text.Bounds().H

			ctx.DrawText(0, y, text)
		}
	}

	c.RenderTo(pdf)

	pdf.NewPage(width, height)
}

func findFontSize(font font.Font, input string, rule config.Render) float64 {
	lastValidSize := 0.0

	for size := 1.0; ; size += 1.0 {
		face := font.Font.Face(size, canvas.Black)
		face.Font.SetFeatures(rule.Rules.Features)

		txt := canvas.NewTextBox(face, input, rule.Rules.Width, 0.0, canvas.Left, canvas.Top, 0.0, 0.0)

		if txt.Bounds().W > rule.Rules.Width-10 || txt.Bounds().H >= rule.Rules.Height-10 {
			return lastValidSize
		}

		lastValidSize = size
	}
}
