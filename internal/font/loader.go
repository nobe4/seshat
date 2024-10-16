package font

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/tdewolff/canvas"
)

type Font struct {
	Name  string
	Path  string
	Font  *canvas.Font
	Order int
}

type Fonts []Font

var (
	fontOrder = map[string]int{
		"thin":       0,
		"extralight": 1,
		"light":      2,
		"regular":    3,
		"medium":     4,
		"extrabold":  5,
		"bold":       6,
		"black":      7,
		"heavy":      8,
	}
)

func LoadFromDir(dir string) (Fonts, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	fonts := make(Fonts, 0, len(files))

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		fPath := path.Join(dir, f.Name())

		font, err := canvas.LoadFontFile(fPath, canvas.FontRegular)
		if err != nil {
			fmt.Printf("error loading font %s: %v\n", fPath, err)
			continue
		}

		nameWithoutExt := strings.TrimSuffix(f.Name(), path.Ext(f.Name()))
		lowerName := strings.ToLower(nameWithoutExt)

		order := -1

		for n, o := range fontOrder {
			if strings.Contains(lowerName, n) {
				order = o
			}
		}

		fonts = append(fonts, Font{
			Name:  f.Name(),
			Path:  fPath,
			Font:  font,
			Order: order,
		})
	}

	if len(fonts) == 0 {
		return nil, fmt.Errorf("no fonts found in %s", dir)
	}

	fonts.Sort()

	return fonts, nil
}

func (f Fonts) Sort() {
	sort.Slice(f, func(i, j int) bool {
		return f[i].Order < f[j].Order
	})
}
