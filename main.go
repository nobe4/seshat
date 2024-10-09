package main

import (
	"github.com/go-pdf/fpdf"
)

func main() {
	pdf := fpdf.New("P", "mm", "A4", ".")
	pdf.AddUTF8Font("f", "B", "./font.ttf")

	pdf.AddPage()
	for i := 40; i > 10; i -= 5 {
		pdf.SetFont("f", "B", float64(i))
		_, unit := pdf.GetFontSize()
		pdf.Write(unit, makeStr("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", 10))
	}

	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		panic(err)
	}
}

func makeStr(s string, c int) string {
	r := ""
	for i := 0; i < c; i++ {
		r += s
	}
	return r
}
