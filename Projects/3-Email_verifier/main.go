package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("domain, hasMX, hasSPR, hasDMARC, sprRecord, dmarcRecord \n")
	//MX -> Mail Exchange: specify the mail servers responsible for receiving email messages on behalf of a domain.
	//SPF -> Sender Policy Framework: SPF is a DNS TXT record that specifies which mail servers are authorized to send emails on behalf of a domain.
	//DMARC -> Domain-based Message Authentication, Reporting, and Conformance: provide domain owners with more control over how their emails are authenticated and how unauthenticated emails should be treated.
	
	fmt.Println("Enter the domain")

	input,err := scanner.ReadString('\n')
	domain := strings.TrimSpace(input)
		checkDomain(domain)
	if err != nil {
			log.Fatal("Error while reading domain and error :  %v", err)
	}
	
}
	
func checkDomain(domain string) {
	var hasMX, hasSPR, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords , err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error in mxRecords and error %v: ",err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords , err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error in txtRecords and error %v: ",err)
	}
	for _,record := range txtRecords{
		if strings.HasPrefix(record,"v=spf1") {
			hasSPR = true
			spfRecord = record
			break
		}
	}
	
	dmarcRecords , err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error in dmarcRecords and error %v: ",err)
	}
	for _,record := range dmarcRecords{
		if strings.HasPrefix(record,"v=DMARC1") {
			hasDMARC = true
			spfRecord = record
			break
		}
	}
	fmt.Printf("%v,%v,%v,%v,%v,%v", domain,hasMX,hasSPR,hasDMARC, spfRecord,dmarcRecord)
}