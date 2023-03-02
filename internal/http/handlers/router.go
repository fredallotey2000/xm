package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ErrorBadRequest = errors.New("invalid request body")

type ErrorResp struct {
	Error string `json:"error,omitempty"`
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

//writeResponse writes the response to the http reponse object
func writeResponse(w http.ResponseWriter, status int, data interface{}, err error) {
	if data == nil {
		w.WriteHeader(status)
		fmt.Fprint(w)
		return
	}
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "No record found")
			return
		}
		resp := ErrorResp{
			Error: fmt.Sprint(err),
		}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			fmt.Fprintf(w, "error encoding resp %v:%v", resp, err)
			return
		}
	}
	w.WriteHeader(status)
	if status == http.StatusCreated {
		s, _ := data.(string)
		w.Header().Add("location", s)
		fmt.Fprint(w)
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Fprintf(w, "error encoding resp %v:%v", data, err)
		return
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
