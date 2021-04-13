package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

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
	oid_sys_model    string = "1.3.6.1.2.1.1.1.0"
)

func main() {
	csvfile, err := os.Open("Printer.csv")
	if err != nil {
		printQuit(err)
	}

	r := csv.NewReader(csvfile)

	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			printQuit(err)
		}
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
}

func printQuit(e error) {
	fmt.Println(e)
	os.Exit(1)
}

func isLow(v float32) string {
	if v < 20 {
		return "\u25c0"
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
	gosnmp.Default.Target = host
	if err := gosnmp.Default.Connect(); err != nil {
		printQuit(err)
	}
	defer gosnmp.Default.Conn.Close()

	retVal, err := gosnmp.Default.Get([]string{oid})
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
