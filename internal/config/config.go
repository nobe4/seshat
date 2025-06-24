package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const warning = `
/!\ README /!\
Could not find config file '%s'.
Make sure the file exists and is readable.
See above for the tried paths.
See https://github.com/nobe4/seshat/blob/main/config.yaml for an example config file.

`

type Config struct {
	Dir  string `yaml:"dir"`
	Path string `yaml:"path"`

	Font   string `yaml:"font"`
	Output string `yaml:"output"`

	Rules   Rules    `yaml:"rules"`
	Renders []Render `yaml:"renders"`
}

type Render struct {
	Type     string   `yaml:"type"`
	Features string   `yaml:"features"`
	Inputs   []string `yaml:"inputs"`
	Rules    Rules    `yaml:"rules"`
}

type Rules struct {
	// Common
	Width      float64   `yaml:"width"`
	Height     float64   `yaml:"height"`
	Size       float64   `yaml:"size"`
	Margins    []float64 `yaml:"margins"`
	Features   string    `yaml:"features"`
	Responsive bool      `yaml:"responsive"`

	// Only for grid
	Columns int `yaml:"columns"`
}

func Read(p string) (Config, error) {
	c := Config{}

	p, content, err := findConfig(p)
	if err != nil {
		return c, err
	}

	fmt.Printf("Found config at %s\n", p)
	c.Dir = filepath.Dir(p)
	c.Path = p

	if err := yaml.Unmarshal(content, &c); err != nil {
		fmt.Printf("Failed to unmarshall the configuration file at '%s'."+
			"Make sure it's a valid YAML file and refer to the documentation for the expected format.\n"+
			"See error below:\n"+
			"%v\n", p, err)
		return c, err
	}

	c.Output = path.Join(path.Dir(c.Path), c.Output)
	c.Font = path.Join(path.Dir(c.Path), c.Font)

	c.PropagateDefaults()
	c.computeMargins()

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

	fullPath, err := filepath.Abs(path)
	content, err = readConfig(fullPath)
	if err == nil {
		return fullPath, content, nil
	}

	fmt.Printf(warning, path)
	return "", nil, fmt.Errorf("could not find config file from path %s", path)
}

func readConfig(path string) ([]byte, error) {
	fmt.Printf("Reading config from %s\n", path)

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read the config file at '%s': %w", path, err)
	}

	return content, nil
}

func processDir() string {
	wd, err := os.Getwd()

	if err != nil {
		fmt.Printf("Failed to get the current working directory: %v\n", err)
		return "."
	}

	return wd
}

func execDir() string {
	cwd, err := os.Executable()

	if err != nil {
		fmt.Printf("Failed to get the executable directory: %v\n", err)
		return "."
	}

	return filepath.Dir(cwd)
}

func (c *Config) PropagateDefaults() {
	for i := range c.Renders {
		c.Renders[i].PropagateDefaults(c.Rules)
	}
}

func (r *Render) PropagateDefaults(defaults Rules) {
	if r.Rules.Width == 0 {
		r.Rules.Width = defaults.Width
	}

	if r.Rules.Height == 0 {
		r.Rules.Height = defaults.Height
	}

	if r.Rules.Size == 0 {
		r.Rules.Size = defaults.Size
	}

	if len(r.Rules.Margins) == 0 {
		r.Rules.Margins = defaults.Margins
	}

	if r.Rules.Features == "" && defaults.Features != "none" {
		r.Rules.Features = defaults.Features
	} else if r.Rules.Features == "none" {
		r.Rules.Features = ""
	}

	if r.Rules.Columns == 0 {
		r.Rules.Columns = defaults.Columns
	}
}

func (c *Config) computeMargins() {
	for i := range c.Renders {
		fmt.Printf("Margins %d: %v\n", i, c.Renders[i].Rules.Margins)
		margins := c.Renders[i].Rules.Margins
		switch len(margins) {
		case 0:
			c.Renders[i].Rules.Margins = []float64{0, 0, 0, 0}
		case 1:
			c.Renders[i].Rules.Margins = []float64{margins[0], margins[0], margins[0], margins[0]}
		case 2:
			c.Renders[i].Rules.Margins = []float64{margins[0], margins[1], margins[0], margins[1]}
		case 3:
			c.Renders[i].Rules.Margins = []float64{margins[0], margins[1], margins[2], margins[1]}
		case 4:
			// do nothing
		default:
			c.Rules.Margins = margins[:4]
		}
		fmt.Printf("Margins %d: %v\n", i, c.Renders[i].Rules.Margins)
	}
}

func (c Config) String() string {
	out, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Sprintf("Failed to marshal the configuration: %v\n", err)
	}

	return string(out)
}
