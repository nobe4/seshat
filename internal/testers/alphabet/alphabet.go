package alphabet

import (
	"test/internal/font"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
)

const (
	alphabet = `a b c d e f g h i j k l m n o p q r s t u v w x y z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 0 1 2 3 4 5 6 7 8 9`
)

func Test(pdf *pdf.PDF, fonts font.Fonts) {
	width, height := pdf.Size()

	for _, font := range fonts {
		c := canvas.New(width, height)
		ctx := canvas.NewContext(c)

		size := 30.0
		for y, i := height, 30; y > 0 && i > 0; i-- {
			face := font.Font.Face(size, canvas.Black)

			txt := canvas.NewTextBox(face, alphabet, width, 0.0, canvas.Left, canvas.Top, 0.0, 0.0)
			ctx.DrawText(0, y, txt)

			y -= txt.Bounds().H
			size -= 1.6
		}

		c.RenderTo(pdf)

		pdf.NewPage(width, height)
	}
}
