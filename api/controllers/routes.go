package controllers

import "github.com/MD-ARMAN-Shanto/gostack/api/middlewares"

func (server *Server) initializeRoutes() {

	// Home routes
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJson(server.Home)).Methods("GET")

	// Login Route
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJson(server.Login)).Methods("POST")

	// user Routes
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJson(server.CreateUser)).Methods("POST")
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJson(server.GetUsers)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJson(server.GetUser)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJson(middlewares.SetMiddlewareAuthentication(server.UpdateUser))).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJson(server.DeleteUser)).Methods("DELETE")

	// post Routes
	server.Router.HandleFunc("/posts", middlewares.SetMiddlewareJson(server.CreatePost)).Methods("POST")
	server.Router.HandleFunc("/posts", middlewares.SetMiddlewareJson(server.GetPosts)).Methods("GET")
	server.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJson(server.GetPost)).Methods("GET")
	server.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJson(middlewares.SetMiddlewareAuthentication(server.UpdatePost))).Methods("PUT")
	server.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJson(server.DeletePost)).Methods("DELETE")
}
