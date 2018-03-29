package main

import (
	. "github.com/ory/ladon"
)

// A bunch of exemplary policies
var Polices = []Policy{
	&DefaultPolicy{
		ID: "1",
		Description: `This policy allows max, peter, zac and ken to create, delete and get the listed resources,
			but only if the client ip matches and the request states that they are the owner of those resources as well.`,
		Subjects:  []string{"max", "peter", "<zac|ken>"},
		Resources: []string{"qrn:some.domain.com:resource:123", "qrn:some.domain.com:resource:345", "qrn:something:foo:<.*>"},
		Actions:   []string{"<create|delete>", "get"},
		Effect:    AllowAccess,
		Conditions: Conditions{
			"owner": &EqualsSubjectCondition{},
			"clientIP": &CIDRCondition{
				CIDR: "127.0.0.1/32",
			},
		},
	},
	&DefaultPolicy{
		ID:          "2",
		Description: "This policy allows max to put object to QingStor bucket 'max'",
		Subjects:    []string{"users:max"},
		Actions:     []string{"qstor:PutObject"},
		Resources:   []string{"qrn:qcs:qstor:::max/<.*>"},
		Effect:      AllowAccess,
	},
	&DefaultPolicy{
		ID:          "3",
		Description: "This policy denies max to put object to object name 'min' on bucket 'max'",
		Subjects:    []string{"users:max"},
		Actions:     []string{"qstor:PutObject"},
		Resources:   []string{"qrn:qcs:qstor:::max/min"},
		Effect:      DenyAccess,
	},
	&DefaultPolicy{
		ID:          "4",
		Description: "This policy allows lucas to perform actions STS:* on any of the resources",
		Subjects:    []string{"users:lucas"},
		Actions:     []string{"STS:<.*>"},
		Resources:   []string{"<.*>"},
		Effect:      AllowAccess,
	},
	&DefaultPolicy{
		ID:          "5",
		Description: "This policy revoke all the access requests from lucas",
		Subjects:    []string{"users:lucas"},
		Actions:     []string{"<.*>"},
		Resources:   []string{"<.*>"},
		Effect:      DenyAccess,
	},
}
