package config

import (
	"fmt"

	"github.com/nobe4/seshat/internal/font"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Rules []Rule `yaml:"rules"`
}

type Rule struct {
	Type string   `yaml:"type"`
	Args []string `yaml:"args"`
}

var DefaultConfig = Config{
	Rules: []Rule{
		{
			Type: "text",
			Args: []string{
				"The quick brown fox jumps over the lazy dog",
				"Sphinx of black quartz, judge my vow",
			},
		},
		{
			Type: "grid",
			Args: []string{
				"a",
				"b",
				"c",
				"test",
			},
		},
		{Type: "alphabet"},
		{Type: "lorem"},
	},
}

func Render(c Config, pdf *pdf.PDF, fonts font.Fonts) {
	out, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
	face := fonts[0].Font.Face(10.0, canvas.Black)

	width, height := pdf.Size()

	can := canvas.New(width, height)
	ctx := canvas.NewContext(can)

	txt := canvas.NewTextLine(face, string(out), canvas.Left)
	ctx.DrawText(5, height-5, txt)

	can.RenderTo(pdf)
}
