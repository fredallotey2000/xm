package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ErrorBadRequest = errors.New("invalid request body")

type ErrorResp struct {
	Error string
}

// Interface to be implemented by all routers
type Router interface {
	Get(uri string, f func(w http.ResponseWriter, r *http.Request))
	Post(uri string, f func(w http.ResponseWriter, r *http.Request))
	Put(uri string, f func(w http.ResponseWriter, r *http.Request))
	Patch(uri string, f func(w http.ResponseWriter, r *http.Request))
	Delete(uri string, f func(w http.ResponseWriter, r *http.Request))

	Start()
	Stop()
}

// writeResponse writes the response to the http reponse object
func WriteResponse(w http.ResponseWriter, status int, data interface{}, err error) {
	w.WriteHeader(status)
	if err != nil {
		resp := ErrorResp{
			Error: fmt.Sprint(err),
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			fmt.Fprintf(w, "error encoding resp %v:%s", resp, err)
		}
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Fprintf(w, "error encoding resp %v:%s", data, err)
	}
}
