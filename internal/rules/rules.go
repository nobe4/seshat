package rules

import (
	"fmt"

	"github.com/nobe4/seshat/internal/font"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
	"gopkg.in/yaml.v3"
)

type Rule struct {
	Test string   `yaml:"test"`
	Args []string `yaml:"args"`
}

var DefaultRules = []Rule{
	{
		Test: "text",
		Args: []string{
			"The quick brown fox jumps over the lazy dog",
			"Sphinx of black quartz, judge my vow",
		},
	},
	{
		Test: "grid",
		Args: []string{
			"a",
			"b",
			"c",
			"test",
		},
	},
	{Test: "alphabet"},
	{Test: "lorem"},
}

func Render(rules []Rule, pdf *pdf.PDF, fonts font.Fonts) {
	out, err := yaml.Marshal(rules)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
	face := fonts[0].Font.Face(10.0, canvas.Black)

	width, height := pdf.Size()

	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	txt := canvas.NewTextLine(face, string(out), canvas.Left)
	ctx.DrawText(5, height-5, txt)

	c.RenderTo(pdf)
}
