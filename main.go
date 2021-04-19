package main

import (
	"flag"
	"io"
	"sync"
)

func main() {

	filePtr := flag.String("f", "printer.csv", "Path to csv printer file.")
	flag.Parse()

	var wg sync.WaitGroup
	r := setupCSV(*filePtr)

	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			printQuit(err)
		}
		go poll(&wg, record)
		wg.Add(1)
	}
	wg.Wait()
}
