package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

func getRecord(svc *route53.Route53) string {
	params := &route53.ListResourceRecordSetsInput{
		HostedZoneId:    aws.String(zoneID),
		StartRecordName: aws.String(name),
	}

	rrSet, _ := svc.ListResourceRecordSets(params)

	rr := rrSet.ResourceRecordSets[0]
	ipaddr := *rr.ResourceRecords[0].Value

	return ipaddr
}
