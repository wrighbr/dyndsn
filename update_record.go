package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

func updateRecord(svc *route53.Route53, recordname string, ipaddr string, zoneID string) {
	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{ // Required
			Changes: []*route53.Change{ // Required
				{ // Required
					Action: aws.String("UPSERT"), // Required
					ResourceRecordSet: &route53.ResourceRecordSet{ // Required
						Name: aws.String(recordname), // Required
						Type: aws.String("A"),        // Required
						ResourceRecords: []*route53.ResourceRecord{
							{ // Required
								Value: aws.String(ipaddr), // Required
							},
						},
						TTL: aws.Int64(60),
					},
				},
			},
		},
		HostedZoneId: aws.String(zoneID), // Required
	}
	resp, err := svc.ChangeResourceRecordSets(params)
	if err != nil {

	}
	fmt.Println(resp)
}

// var ipaddr string
