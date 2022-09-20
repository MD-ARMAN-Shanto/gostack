package controllers

import (
	"github.com/MD-ARMAN-Shanto/gostack/api/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to the Awesome API")
}
