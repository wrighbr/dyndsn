package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

var (
	zoneID string
	name   string
)

func init() {
	flag.StringVar(&name, "n", "", "DNS Record name")
	flag.StringVar(&zoneID, "z", "", "Zone ID")

}

func main() {
	flag.Parse()

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println(err.Error())
	}
	svc := route53.New(sess)

	ipaddr := getIP()
	dnsrecord := getRecord(svc)

	fmt.Println("Current IP Address:", ipaddr)
	fmt.Println("DNS Record IP:", dnsrecord)

	if ipaddr != dnsrecord {
		fmt.Println("IP Address has changed. Now updating DNS Recorded.")
		updateRecord(svc, name, ipaddr, zoneID)
	} else {
		fmt.Println("Current IP and DNS Record are the same. No need to change the record!")
	}

}
