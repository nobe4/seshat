package main

import (
	"fmt"
	"os"
	"test/internal/font"
	"test/internal/testers/lettergrid"
	"time"

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

	// alphabet.Test(pdf, fonts)
	// lorem.Test(pdf, fonts)
	lettergrid.Test(pdf, fonts)

	if err := pdf.Close(); err != nil {
		panic(err)
	}

	end := time.Now()
	fmt.Printf("Ran at %s in %fs\n", end.Format("15:04:05"), end.Sub(start).Seconds())
}
