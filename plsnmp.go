package main

import (
	"fmt"
	"time"

	"github.com/gosnmp/gosnmp"
)

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
		if oid[1:] == "" {
			isOID, retVal = false, "0"
		} else {
			isOID, retVal = false, oid[1:]
		}
	} else {
		isOID, retVal = true, oid
	}
	return
}
