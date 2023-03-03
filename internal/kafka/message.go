package kafka

import (
	cmp "xm/pkg/company"
)

type Message struct {
	cmp.Company
	cmp.CompanyPatch
	Id            string
	RequestMethod string
}
