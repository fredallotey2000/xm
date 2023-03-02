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

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

//writeResponse writes the response to the http reponse object
func writeResponse(w http.ResponseWriter, status int, data interface{}, err error) {
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

// // writeResponse is a helper method that allows to write and HTTP status & response
// func writeResponse2(w http.ResponseWriter, status int, data interface{}, err error) {
// 	w.WriteHeader(status)
// 	resp := Response{
// 		Data: data,
// 	}
// 	if err != nil {
// 		resp.Error = fmt.Sprint(err)
// 	}
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	if err := json.NewEncoder(w).Encode(resp); err != nil {
// 		fmt.Fprintf(w, "error encoding resp %v:%s", resp, err)
// 	}
// }
