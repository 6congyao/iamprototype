package polices

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
		Resources: []string{"myrn:some.domain.com:resource:123", "myrn:some.domain.com:resource:345", "myrn:something:foo:<.+>"},
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
		Description: "This policy allows max to update any resource",
		Subjects:    []string{"max"},
		Actions:     []string{"update"},
		Resources:   []string{"<.*>"},
		Effect:      AllowAccess,
	},
	&DefaultPolicy{
		ID:          "3",
		Description: "This policy denies max to broadcast any of the resources",
		Subjects:    []string{"max"},
		Actions:     []string{"broadcast"},
		Resources:   []string{"<.*>"},
		Effect:      DenyAccess,
	},
	&DefaultPolicy{
		ID:          "4",
		Description: "This policy denies max to broadcast any of the resources",
		Subjects:    []string{"users:lucas"},
		Actions:     []string{"STS:<.*>"},
		Resources:   []string{"<.*>"},
		Effect:      AllowAccess,
	},
	//&DefaultPolicy{
	//	ID:          "5",
	//	Description: "This policy denies max to broadcast any of the resources",
	//	Subjects:    []string{"users:lucas"},
	//	Actions:     []string{"STS:AssumeRole"},
	//	Resources:   []string{"<.*>"},
	//	Effect:      DenyAccess,
	//},
}
