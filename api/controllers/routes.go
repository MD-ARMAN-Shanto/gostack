package controllers

import (
	"github.com/MD-ARMAN-Shanto/gostack/api/middlewares"
)

func (s *Server) initializeRoutes() {

	// Home routes
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJson(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJson(s.Login)).Methods("POST")

	// user Routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJson(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJson(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJson(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJson(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJson(s.DeleteUser)).Methods("DELETE")

	// post Routes
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJson(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJson(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJson(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJson(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJson(s.DeletePost)).Methods("DELETE")
}
