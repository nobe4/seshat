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

	run(config)
}

func run(c config.Config) {
	start := time.Now()
	fmt.Printf("Start at %s\n", start.Format("15:04:05"))

	f, err := os.Create(c.Output)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pdf := pdf.New(f, width, height, &pdf.Options{})

	fonts, err := font.LoadFromDir(c.Font)
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
