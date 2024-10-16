package config

import (
	"fmt"
	"os"
	"path"

	"github.com/nobe4/seshat/internal/font"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Font   string `yaml:"font"`
	Output string `yaml:"output"`

	Rules []Rule `yaml:"rules"`
}

type Rule struct {
	Type string   `yaml:"type"`
	Args []string `yaml:"args"`
}

var DefaultConfig = Config{
	Font:   ".",
	Output: "output.pdf",
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

func Read(p string) Config {
	content, err := os.ReadFile(p)
	if err != nil {
		fmt.Printf("Error reading %s: %w\n", p, err)
		return DefaultConfig
	}

	c := DefaultConfig
	if err := yaml.Unmarshal(content, &c); err != nil {
		fmt.Printf("Error unmarshalling %s: %w\n", p, err)
		return DefaultConfig
	}

	c.Output = path.Join(path.Dir(p), c.Output)
	c.Font = path.Join(path.Dir(p), c.Font)

	return c
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
