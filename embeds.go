package main

import (
	"embed"
	"encoding/csv"
	"errors"
	"io/fs"
	"os"
)

//go:embed Printer.csv
var printers embed.FS

type printerCol int

const (
	descr printerCol = iota
	host             = iota
	level            = iota
	cap              = iota
	model            = iota
)

func setupCSV(lf string) (csvReader *csv.Reader) {

	var f fs.File
	var err error

	f, err = os.Open(lf)

	if err != nil {
		f, err = printers.Open("Printer.csv")
		if err != nil {
			printQuit(errors.New("no available printer file"))
		}
	}

	csvReader = csv.NewReader(f)
	csvReader.Comment = '#'

	return
}
