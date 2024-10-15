package main

import (
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
	outFile = "test.pdf"

	// pdf size in points
	width  = 210.0
	height = 297.0
)

func main() {
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

	r := getRules()

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

func getRules() []rules.Rule {
	c, err := os.ReadFile("rules.yaml")
	if err != nil {
		fmt.Println("Error reading rules.yaml:", err)
		return rules.DefaultRules
	}

	var r []rules.Rule
	if err := yaml.Unmarshal(c, &r); err != nil {
		fmt.Println("Error unmarshalling rules.yaml:", err)
		return rules.DefaultRules
	}

	return r
}
