package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gosnmp/gosnmp"
)

const (
	descr int = iota
	host  int = iota
	level int = iota
	cap   int = iota
	model int = iota
)

const (
	oid_sys_location string = "1.3.6.1.2.1.1.6.0"
	oid_sys_name     string = "1.3.6.1.2.1.1.5.0"
	oid_sys_contact  string = "1.3.6.1.2.1.1.4.0"
)

func main() {
	r := setupCSV("Printer.csv")

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
func setupCSV(filename string) *csv.Reader {
	csvfile, err := os.Open(filename)
	if err != nil {
		printQuit(err)
	}

	return csv.NewReader(csvfile)
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
		return "***"
	}
	return ""
}

func getSNMPInt(host string, oid string) int {
	return int(gosnmp.ToBigInt(getSNMPValue(host, oid)).Int64())
}

func getSNMPString(host string, oid string) string {
	return fmt.Sprintf("%s", getSNMPValue(host, oid))
}

func getSNMPValue(host string, oid string) interface{} {

	params := &gosnmp.GoSNMP{
		Target:    host,
		Port:      161,
		Community: "public",
		Version:   gosnmp.Version2c,
		//Logger:    log.New(os.Stdout, "", 0),
		Timeout: time.Duration(2) * time.Second,
	}

	if err := params.Connect(); err != nil {
		printQuit(err)
	}
	defer params.Conn.Close()

	retVal, err := params.Get([]string{oid})
	if err != nil {
		printQuit(err)
	}
	return retVal.Variables[0].Value

}

func parseOID(oid string) (isOID bool, retVal string) {
	if oid[0] == 'M' {
		isOID, retVal = false, oid[1:]
	} else {
		isOID, retVal = true, oid
	}
	return
}
