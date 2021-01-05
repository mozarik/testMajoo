package controllers

import "github.com/mozarik/testMajoo/api_model/middlewares"

func (s *Server) initializeRoutes() {

	// Home page
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// login
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//user route
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

}
