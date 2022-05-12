package controllers

import "github.com/mikorail/go-api/api/middlewares"

func (s *Server) initializeRoutes() {
	//Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Auth Route
	s.Router.HandleFunc("/api/auth/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/api/auth/register", middlewares.SetMiddlewareJSON(s.Register)).Methods("POST")

}
