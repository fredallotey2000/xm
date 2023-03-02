package handlers

import (
	//"encoding/json"
	"encoding/json"
	"net/http"

	v "xm/internal/http/validator"
	l "xm/internal/logger"
	usr "xm/pkg/user"

	mw "xm/internal/http/middlewares"
)

// handler holds all the dependencies required for server requests
type userHandler struct {
	service   usr.Service
	validator v.Validator
	lg        l.Logger
}

type Token struct {
	Token string `json:"token"`
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
	var user *usr.User

	// Read the request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.lg.HttpError(r, "AuthenticateUser", err)
		writeResponse(w, http.StatusBadRequest, nil, ErrorBadRequest)
		return
	}
	//Validates the JSON object and makes sure it meets the required request fields
	if err := h.validator.ValidateJSON(user); err != nil {
		h.lg.HttpError(r, "AuthenticateUser", err)
		writeResponse(w, http.StatusBadRequest, nil, ErrorBadRequest)
		return
	}

	//Get the Databank location from the service layer
	user, err := h.service.AuthUser(r.Context(), user.Email, user.Password)
	if err != nil {
		h.lg.HttpError(r, "AuthenticateUser", err)
		writeResponse(w, http.StatusInternalServerError, nil, err)
		return
	}

	jwtToken, err := mw.GenerateJWT(user.Email, user.Password)
	token := Token{
		Token: jwtToken,
	}
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	writeResponse(w, http.StatusOK, token, nil)

}
