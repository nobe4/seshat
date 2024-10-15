package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nobe4/seshat/internal/font"
	"github.com/nobe4/seshat/internal/rules"
	"github.com/nobe4/seshat/internal/testers"
	"github.com/tdewolff/canvas/renderers/pdf"
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

	r := rules.DefaultRules
	rules.Render(r, pdf, fonts)
	for _, rule := range r {
		fmt.Printf("Running rule %s(%v)\n", rule.Test, rule.Args)
		testers.Get(rule.Test)(pdf, fonts, rule.Args)
	}

	if err := pdf.Close(); err != nil {
		panic(err)
	}

	end := time.Now()
	fmt.Printf("Ran at %s in %fs\n", end.Format("15:04:05"), end.Sub(start).Seconds())
}
