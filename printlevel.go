package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

type printerCol int

const (
	descr printerCol = iota
	host             = iota
	level            = iota
	cap              = iota
	model            = iota
)

const (
	oid_sys_location string = "1.3.6.1.2.1.1.6.0"
	oid_sys_name     string = "1.3.6.1.2.1.1.5.0"
	oid_sys_contact  string = "1.3.6.1.2.1.1.4.0"
	oid_sys_model    string = "1.3.6.1.2.1.1.1.0"
)

func main() {

	filePtr := flag.String("f", "printer.csv", "Path to csv printer file.")
	flag.Parse()

	r := setupCSV(*filePtr)
	var wg sync.WaitGroup

	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			printQuit(err)
		}
		go run(&wg, record)
		wg.Add(1)
	}
	wg.Wait()
}

func setupCSV(filename string) (csvReader *csv.Reader) {
	csvfile, err := os.Open(filename)
	if err != nil {
		printQuit(err)
	}

	csvReader = csv.NewReader(csvfile)
	csvReader.Comment = '#'

	return
}

func run(wg *sync.WaitGroup, record []string) {
	defer wg.Done()
	//	fmt.Println(record)
	var err error
	var printLevel int
	if isOID, OID := parseOID(record[level]); isOID {
		printLevel = getSNMPInt(record[host], OID)
	} else {
		printLevel, err = strconv.Atoi(OID)
		if err != nil {
			printQuit(errors.New("bad manual value"))
		}
	}

	var printCap int
	if isOID, OID := parseOID(record[cap]); isOID {
		printCap = getSNMPInt(record[host], OID)
	} else {
		printCap, err = strconv.Atoi(OID)
		if err != nil {
			printQuit(errors.New("bad manual value"))
		}
	}
	levelPercent := float32(printLevel) / float32(printCap) * 100

	fmt.Printf("%-24v %-24s (%d/%d) %.2f%%%v\n",
		getSNMPString(record[host], oid_sys_name),
		record[descr],
		printLevel,
		printCap,
		levelPercent,
		isLow(levelPercent))

}

func printQuit(e error) {
	fmt.Println(e)
	os.Exit(1)
}

func isLow(v float32) string {
	if v < 20 {
		return "\u25c0\u25c0\u25c0"
	}
	return ""
}
