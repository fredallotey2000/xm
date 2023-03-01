package handlers

import (
	//"encoding/json"
	"encoding/json"
	"net/http"

	v "xm/internal/http/validator"
	l "xm/internal/logger"
	usr "xm/pkg/user"
)

// handler holds all the dependencies required for server requests
type userHandler struct {
	service   usr.Service
	validator v.Validator
	lg        l.Logger
}

// Handler is the interface we expose to outside packages
type UserHandler interface {
	AuthenticateUser(w http.ResponseWriter, r *http.Request)
}

//Creating a new Handler
func NewUserHandler(ds usr.Service, v v.Validator, log l.Logger) UserHandler {
	// Create a new handler with a logger, a repo and a validator
	h := &userHandler{
		service:   ds,
		validator: v,
		lg:        log,
	}
	return h
}

/*LocateDataBank writes a JSON response to the http interface, by
making a request to a repository to get the location of a databank*/
func (h *userHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var user usr.User

	// Read the request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.lg.HttpError(r, "AuthenticateUser", err)
		WriteResponse(w, http.StatusBadRequest, nil, ErrorBadRequest)
		return
	}
	//Validates the JSON object and makes sure it meets the required request fields
	if b, err := h.validator.ValidateJSON(user); !b {
		h.lg.HttpError(r, "AuthenticateUser", err)
		WriteResponse(w, http.StatusBadRequest, nil, ErrorBadRequest)
		return
	}

	//Get the Databank location from the service layer
	hops, err := h.service.AuthUser(r.Context(), user.Email, user.Password)
	if err != nil {
		h.lg.HttpError(r, "CreateCompany", err)
		WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	//write response wiht datbank location to http response
	WriteResponse(w, http.StatusOK, hops, nil)
}
