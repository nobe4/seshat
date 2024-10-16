package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/nobe4/seshat/internal/font"
	"github.com/nobe4/seshat/internal/rules"
	"github.com/nobe4/seshat/internal/testers"
	"github.com/tdewolff/canvas/renderers/pdf"
	"gopkg.in/yaml.v3"
)

const (
	fontDir = "./font/OTF"

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

func run(r []rules.Rule, outFile, fontDir string) {
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

	rules.Render(r, pdf, fonts)
	for _, rule := range r {
		fmt.Printf("Running rule %s(%v)\n", rule.Type, rule.Args)
		testers.Get(rule.Type)(pdf, fonts, rule.Args)
	}

	if err := pdf.Close(); err != nil {
		panic(err)
	}

	end := time.Now()
	fmt.Printf("Ran at %s in %fs\n", end.Format("15:04:05"), end.Sub(start).Seconds())
}

func getConfig(p string) []rules.Rule {
	c, err := os.ReadFile(p)
	if err != nil {
		fmt.Printf("Error reading %s: %w\n", p, err)
		return rules.DefaultRules
	}

	var r []rules.Rule
	if err := yaml.Unmarshal(c, &r); err != nil {
		fmt.Printf("Error unmarshalling %s: %w\n", p, err)
		return rules.DefaultRules
	}

	return r
}
