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
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	post.Prepare()
	err = post.Validate()
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	if uid != post.AuthorId {
		responses.Error(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	postCreated, err := post.PostSave(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.Error(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, postCreated.ID))
	responses.JSON(w, http.StatusCreated, postCreated)
}

func (server *Server) GetPosts(w http.ResponseWriter, r *http.Request) {

	post := models.Post{}

	posts, err := post.FindAllPosts(server.DB)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

func (server *Server) GetPost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}

	postReceived, err := post.FindPostById(server.DB, pid)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, postReceived)
}

func (server *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//check the post id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	//CHeck if the auth token is valid and  get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	//check if the post exist
	post := models.Post{}
	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		responses.Error(w, http.StatusNotFound, errors.New("post not found"))
		return
	}

	// If a user attempt to update a post not belonging to him
	if uid != post.AuthorId {
		responses.Error(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	//read the data posted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	// start processing the request data
	postUpdate := models.Post{}
	err = json.Unmarshal(body, &postUpdate)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	//also check the rquest user id is equal to the one gotten form token
	if uid != postUpdate.AuthorId {
		responses.Error(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	postUpdate.Prepare()
	err = postUpdate.Validate()
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	postUpdate.ID = post.ID

	postUpdated, err := postUpdate.UpdateAPost(server.DB)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, postUpdated)
}

func (server *Server) DeletePost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid post id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	// Check if the post exist
	post := models.Post{}
	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		responses.Error(w, http.StatusNotFound, errors.New("unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this post?
	if uid != post.AuthorId {
		responses.Error(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	_, err = post.DeleteAPost(server.DB, pid, uid)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
