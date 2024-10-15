package testers

import (
	"github.com/nobe4/seshat/internal/font"
	"github.com/nobe4/seshat/internal/testers/alphabet"
	"github.com/nobe4/seshat/internal/testers/grid"
	"github.com/nobe4/seshat/internal/testers/lorem"
	"github.com/nobe4/seshat/internal/testers/text"
	"github.com/tdewolff/canvas/renderers/pdf"
)

type TestFn func(*pdf.PDF, font.Fonts, []string)

func Get(name string) TestFn {
	switch name {
	case "alphabet":
		return alphabet.Test
	case "lorem":
		return lorem.Test
	case "grid":
		return grid.Test
	case "text":
		return text.Test
	}

	panic("undefined tester " + name)
}
