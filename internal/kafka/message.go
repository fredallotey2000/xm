package kafka

import (
	cmp "xm/pkg/company"
)

type Message struct {
	cmp.Company
	Id            string
	RequestMethod string
}
