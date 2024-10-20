package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/nobe4/seshat/internal/config"
	"github.com/nobe4/seshat/internal/font"
	"github.com/nobe4/seshat/internal/testers"
	"github.com/tdewolff/canvas/renderers/pdf"
)

func main() {
	configPtr := flag.String("config", "config.yaml", "path to the configuration file")
	flag.Parse()

	run(*configPtr)
}

func run(configPath string) {
	for {
		fmt.Printf("\n\nRunning with config %s\n", configPath)

		c, err := config.Read(configPath)
		if err != nil {
			fmt.Printf("error reading the config: %v\n", err)
		} else {
			fmt.Println(c)
			configPath = c.Path
		}

		if err := render(c); err != nil {
			fmt.Printf("error rendering: %v\n", err)
		}

		if err := waitForModification(c); err != nil {
			fmt.Printf("error watching for changes: %v\n", err)
		}
	}
}

func waitForModification(c config.Config) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	modifiedChan := make(chan struct{})
	errorChan := make(chan error)

	// TODO: handle signal more gracefully
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					errorChan <- errors.New("watcher closed")
				}

				if event.Has(fsnotify.Write) {
					fmt.Printf("Modified file detected at %s: %s\n",
						time.Now().Format("15:04:05"),
						event.Name)

					modifiedChan <- struct{}{}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					errorChan <- errors.New("watcher closed")
				}

				errorChan <- err
			}
		}
	}()

	fmt.Printf("Watching files in %s\n", c.Dir)
	if err = watcher.Add(c.Dir); err != nil {
		fmt.Printf("error watching files in %s: %v\n", c.Dir, err)
	}

	fmt.Printf("Watching files in %s\n", c.Font)
	if err = watcher.Add(c.Font); err != nil {
		fmt.Printf("error watching files in %s: %v\n", c.Font, err)
	}

	fmt.Println("Waiting for modification, press Ctrl+C to exit")

	select {
	case err := <-errorChan:
		return err
	case <-modifiedChan:
		return nil
	}
}

func render(c config.Config) error {
	start := time.Now()
	fmt.Printf("Start at %s\n", start.Format("15:04:05"))

	f, err := os.Create(c.Output)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer f.Close()

	// TODO: allow to create a pdf without the first page
	pdf := pdf.New(f, c.Defaults.Width, c.Defaults.Height, &pdf.Options{
		SubsetFonts: true,
	})

	fonts, err := font.Load(c.Font)
	if err != nil {
		return fmt.Errorf("error loading fonts: %w", err)
	}

	for _, r := range c.Rules {
		fmt.Printf("Running rule %s(%v)\n", r.Type, r.Args)
		t := testers.Get(r.Type)
		if t != nil {
			t(pdf, fonts, c, r)
		}
	}

	if err := pdf.Close(); err != nil {
		return fmt.Errorf("error closing pdf: %w", err)
	}

	end := time.Now()
	fmt.Printf("Ran at %s in %fs\n", end.Format("15:04:05"), end.Sub(start).Seconds())

	return nil
}
