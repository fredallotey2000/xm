package handlers

import (
	"encoding/json"
	"net/http"

	v "xm/internal/http/validator"
	lgg "xm/internal/logger"
	cmp "xm/pkg/company"

	"github.com/gorilla/mux"

	kf "xm/internal/kafka"
)

// companyHandler holds all the dependencies required for server company requests
type companyHandler struct {
	service      cmp.Service
	validator    v.Validator
	lg           lgg.Logger
	producerChan chan kf.Message
}

// Handler is the interface we expose to outside packages
type CompanyHandler interface {
	CreateCompany(w http.ResponseWriter, r *http.Request)
	UpdateCompany(w http.ResponseWriter, r *http.Request)
	GetCompany(w http.ResponseWriter, r *http.Request)
	DeleteCompany(w http.ResponseWriter, r *http.Request)
	CheckHealth(w http.ResponseWriter, r *http.Request)
}

//Creating a new company Handler
func NewCompanyHandler(pChan chan kf.Message, ds cmp.Service, v v.Validator, log lgg.Logger) CompanyHandler {
	// Create a new handler with a logger, a repo and a validator
	h := &companyHandler{
		service:      ds,
		validator:    v,
		lg:           log,
		producerChan: pChan,
	}
	return h
}

// function to receive and process http posts to create a new company
func (h *companyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var comp cmp.Company

	// Read the request body
	if err := json.NewDecoder(r.Body).Decode(&comp); err != nil {
		h.lg.HttpError(r, "CreateCompany", err)
		writeResponse(w, http.StatusBadRequest, nil, ErrorBadRequest)
		return
	}
	//Validates the JSON object and makes sure it meets the required request fields
	if b, err := h.validator.ValidateJSON(comp); !b {
		h.lg.HttpError(r, "CreateCompany", err)
		writeResponse(w, http.StatusBadRequest, nil, ErrorBadRequest)
		return
	}

	//push onto kafka producer for processing
	kfMsg := kf.Message{
		Company:        comp,
		Request:        r,
		ResponseWriter: w,
	}
	h.producerChan <- kfMsg
}

// function to receive and process http patch to update a new company
func (h *companyHandler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	compId := vars["companyId"]
	var cmp cmp.Company

	// Read the request body
	if err := json.NewDecoder(r.Body).Decode(&cmp); err != nil {
		h.lg.HttpError(r, "UpdateCompany", err)
		writeResponse(w, http.StatusBadRequest, [2]string{}, ErrorBadRequest)
		return
	}
	kfMsg := kf.Message{
		Company:        cmp,
		Id:             compId,
		Request:        r,
		ResponseWriter: w,
	}
	h.producerChan <- kfMsg

}

// function to receive and process http company delete requests
func (h *companyHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	compId := vars["companyId"]

	//push onto kafka producer for processing
	kfMsg := kf.Message{
		Id:             compId,
		Request:        r,
		ResponseWriter: w,
	}

	h.producerChan <- kfMsg

}

// function to handle http get requests to fetch company info using id
func (h *companyHandler) GetCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	compId := vars["companyId"]
	kfMsg := kf.Message{
		Id:             compId,
		Request:        r,
		ResponseWriter: w,
	}
	h.producerChan <- kfMsg
}

// Checks if service is up and running
func (h *companyHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "Healthy", nil)
}
