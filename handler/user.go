package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Hendra-Huang/go-standard-layout/log"
	"github.com/IndraGunawan/gosample"

	"github.com/julienschmidt/httprouter"
)

// UserServicer is a contract of a user related services
type UserServicer interface {
	Create(context.Context, gosample.User) (int64, error)
	FindAll(context.Context) ([]gosample.User, error)
	FindByID(context.Context, int64) (gosample.User, error)
}

// UserHandler handles any http request to user
type UserHandler struct {
	userService UserServicer
}

// NewUserHandler initializes UserHandler instance
func NewUserHandler(userService UserServicer) *UserHandler {
	return &UserHandler{userService}
}

// GetAll handle request then write response that contains all user record
func (uh *UserHandler) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := uh.userService.FindAll(r.Context())
	if err != nil {
		log.Errors(err)
		return
	}

	JSONResponse(w, users)
}

// GetByID handle request then write response that contains single user record by ID
func (uh *UserHandler) GetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	user, err := uh.userService.FindByID(r.Context(), id)

	if err != nil {
		log.Errors(err)
		return
	}

	JSONResponse(w, user)
}

// Create handler request and save new user record
func (uh *UserHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errors(err)
		return
	}

	var user gosample.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Errors(err)
		return
	}

	var lastInsertID int64
	lastInsertID, err = uh.userService.Create(r.Context(), user)
	if err != nil {
		log.Errors(err)
		return
	}

	user.ID = lastInsertID

	JSONResponse(w, user)
}
