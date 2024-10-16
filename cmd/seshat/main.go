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
)

const (
	// pdf size in points
	width  = 210.0
	height = 297.0
)

func main() {
	configPtr := flag.String("config", "config.yaml", "path to the configuration file")
	flag.Parse()

	config := config.Read(*configPtr)

	if err := run(config); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func run(c config.Config) error {
	start := time.Now()
	fmt.Printf("Start at %s\n", start.Format("15:04:05"))

	f, err := os.Create(c.Output)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer f.Close()

	pdf := pdf.New(f, width, height, &pdf.Options{
		SubsetFonts: true,
	})

	fonts, err := font.LoadFromDir(c.Font)
	if err != nil {
		return fmt.Errorf("error loading fonts: %w", err)
	}

	config.Render(c, pdf, fonts)
	for _, rule := range c.Rules {
		fmt.Printf("Running rule %s(%v)\n", rule.Type, rule.Args)
		testers.Get(rule.Type)(pdf, fonts, rule.Args)
	}

	if err := pdf.Close(); err != nil {
		return fmt.Errorf("error closing pdf: %w", err)
	}

	end := time.Now()
	fmt.Printf("Ran at %s in %fs\n", end.Format("15:04:05"), end.Sub(start).Seconds())

	return nil
}
