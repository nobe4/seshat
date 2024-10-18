package testers

import (
	"fmt"

	"github.com/nobe4/seshat/internal/font"
	"github.com/nobe4/seshat/internal/testers/grid"
	"github.com/nobe4/seshat/internal/testers/text"
	"github.com/tdewolff/canvas/renderers/pdf"
)

type TestFn func(*pdf.PDF, font.Fonts, string, []string)

func Get(name string) TestFn {
	switch name {
	case "grid":
		return grid.Test
	case "text":
		return text.Test
	}

	fmt.Printf("Unknown tester %s\n", name)
	return nil
}
