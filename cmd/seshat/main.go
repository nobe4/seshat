package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/nobe4/seshat/internal/config"
	"github.com/nobe4/seshat/internal/font"
	"github.com/nobe4/seshat/internal/testers"
	"github.com/tdewolff/canvas/renderers/pdf"
	"gopkg.in/yaml.v3"
)

const (
	// pdf size in points
	width  = 210.0
	height = 297.0
)

func main() {
	configPtr := flag.String("config", "config.yaml", "path to the configuration file")
	fontPtr := flag.String("font", ".", "path to the font file or directory")
	outputPtr := flag.String("output", "output.pdf", "path to the output file")
	flag.Parse()

	config := getConfig(*configPtr)

	run(config, *outputPtr, *fontPtr)
}

func run(c config.Config, outFile, fontDir string) {
	start := time.Now()
	fmt.Printf("Start at %s\n", start.Format("15:04:05"))

	f, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pdf := pdf.New(f, width, height, &pdf.Options{})

	fonts, err := font.LoadFromDir(fontDir)
	if err != nil {
		panic(err)
	}

	config.Render(c, pdf, fonts)
	for _, rule := range c.Rules {
		fmt.Printf("Running rule %s(%v)\n", rule.Type, rule.Args)
		testers.Get(rule.Type)(pdf, fonts, rule.Args)
	}

	if err := pdf.Close(); err != nil {
		panic(err)
	}

	end := time.Now()
	fmt.Printf("Ran at %s in %fs\n", end.Format("15:04:05"), end.Sub(start).Seconds())
}

func getConfig(p string) config.Config {
	content, err := os.ReadFile(p)
	if err != nil {
		fmt.Printf("Error reading %s: %w\n", p, err)
		return config.DefaultConfig
	}

	var c config.Config
	if err := yaml.Unmarshal(content, &c); err != nil {
		fmt.Printf("Error unmarshalling %s: %w\n", p, err)
		return config.DefaultConfig
	}

	return c
}
