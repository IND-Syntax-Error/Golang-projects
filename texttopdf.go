package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
)

func main() {
	file := "convert.txt" //Change the file name to the one you want to convert
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("%s file not found", file)
	}
	pdf := gofpdf.New("P", "cm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12) //Change the font,style and size as per your wish
	pdf.MultiCell(190, 4, string(content), "0", "0", false)
	_ = pdf.OutputFileAndClose("converted.pdf")//Change the file name of the converted file as per your wish
	fmt.Println("Text has been converted to PDF")
}