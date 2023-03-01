package kafka

import (
	cmp "xm/pkg/company"
	"net/http"
)

type Message struct {
	cmp.Company
	Id             string
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}
