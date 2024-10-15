package lorem

import (
	"github.com/nobe4/seshat/internal/font"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
)

const (
	lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam in dui mauris. Vivamus hendrerit arcu sed`
)

func Test(pdf *pdf.PDF, fonts font.Fonts, _ []string) {
	width, height := pdf.Size()

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
