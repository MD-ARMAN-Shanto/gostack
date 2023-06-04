package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MD-ARMAN-Shanto/gostack/api/auth"
	"github.com/MD-ARMAN-Shanto/gostack/api/models"
	"github.com/MD-ARMAN-Shanto/gostack/api/responses"
	"github.com/MD-ARMAN-Shanto/gostack/api/utils/formaterror"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		responses.Error(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
}

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	userGotten, err := user.FindUserById(server.DB, uint32(uid))
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	if tokenID != uint32(uid) {
		responses.Error(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	}
	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.Error(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUser)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	user := models.User{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	if tokenID != 0 && tokenID != uint32(uid) {
		responses.Error(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	_, err = user.DeleteAUser(server.DB, uint32(uid))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}
