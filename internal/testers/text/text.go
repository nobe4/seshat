package text

import (
	"fmt"
	"strconv"

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
			// TODO: move this in the config parsing
			size := config.Size
			sizeString, ok := rule.Args["size"]
			if ok {
				var err error
				size, err = strconv.ParseFloat(sizeString, 64)
				if err != nil {
					size = config.Size
					fmt.Printf("error parsing size: %v\n", err)
				}
			}

			face := font.Font.Face(size, canvas.Black)
			// TODO: move features in config parsing
			face.Font.SetFeatures(rule.Args["features"])

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
