package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println("domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord")

	for scanner.Scan() {
		domain := scanner.Text()
		checkDomain(domain)
	}
}

func checkDomain(domain string) {
	fmt.Println("----------->", domain)
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecord, err := net.LookupMX(domain)

	if err != nil {
		fmt.Printf("Error:- %v", err.Error())
	}

	if len(mxRecord) > 0 {
		hasMx = true
	}

	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Printf("Error:- %v", err.Error())
	}

	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc" + domain)

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
		}
	}

	fmt.Println(hasSPF, hasMx, hasDMARC, spfRecord, dmarcRecord)
}
