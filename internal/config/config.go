package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/nobe4/seshat/internal/font"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
	"gopkg.in/yaml.v3"
)

const warning = `
/!\ README /!\
Could not find config file '%s'.
Make sure the file exists and is readable.
See above for the tried paths.
See https://github.com/nobe4/seshat/blob/main/examples/config.yaml for an example config file.

`

type Config struct {
	Path   string `yaml:"path"`
	Font   string `yaml:"font"`
	Output string `yaml:"output"`

	Rules []Rule `yaml:"rules"`
}

type Rule struct {
	Type     string   `yaml:"type"`
	Features string   `yaml:"features"`
	Args     []string `yaml:"args"`
}

func Read(p string) (Config, error) {
	c := Config{}

	p, content, err := findConfig(p)
	if err != nil {
		return c, err
	}

	fmt.Printf("Found config at %s\n", p)
	c.Path = p

	if err := yaml.Unmarshal(content, &c); err != nil {
		fmt.Printf("Error unmarshalling %s: %v\n", p, err)
		return c, err
	}

	c.Output = path.Join(path.Dir(c.Path), c.Output)
	c.Font = path.Join(path.Dir(c.Path), c.Font)

	return c, nil
}

func findConfig(path string) (string, []byte, error) {
	processPath := filepath.Join(processDir(), path)
	content, err := readConfig(processPath)
	if err == nil {
		return processPath, content, nil
	}

	execPath := filepath.Join(execDir(), path)
	content, err = readConfig(execPath)
	if err == nil {
		return execPath, content, nil
	}

	fmt.Printf(warning, path)
	return "", nil, fmt.Errorf("could not find config file from path %s", path)
}

func readConfig(path string) ([]byte, error) {
	fmt.Printf("Reading config from %s\n", path)

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading %s: %w", path, err)
	}

	return content, nil
}

func processDir() string {
	wd, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		return "."
	}

	return wd
}

func execDir() string {
	cwd, err := os.Executable()

	if err != nil {
		fmt.Printf("Error getting executable directory: %v\n", err)
		return "."
	}

	return filepath.Dir(cwd)
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
