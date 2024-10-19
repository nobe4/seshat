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

	Width  float64 `yaml:"width"`
	Height float64 `yaml:"height"`
	Size   float64 `yaml:"size"`

	Rules []Rule `yaml:"rules"`
}

type Rule struct {
	Type     string            `yaml:"type"`
	Features string            `yaml:"features"`
	Inputs   []string          `yaml:"inputs"`
	Args     map[string]string `yaml:"args"`
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

func (c Config) String() string {
	out, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}

	return string(out)
}
