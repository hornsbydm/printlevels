package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/gosnmp/gosnmp"
)

func main() {
	csvfile, err := os.Open("Printer.csv")
	if err != nil {
		os.Exit(0)
	}

	r := csv.NewReader(csvfile)

	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}

		gosnmp.Default.Target = record[1]
		err = gosnmp.Default.Connect()
		if err != nil {
			os.Exit(1)
		}
		defer gosnmp.Default.Conn.Close()

		result, err2 := gosnmp.Default.Get([]string{record[2]})

		if err2 != nil {
			os.Exit(2)
		}

		fmt.Printf("%s, %v\n", record[0], result.Variables[0].Value)

	}
}
