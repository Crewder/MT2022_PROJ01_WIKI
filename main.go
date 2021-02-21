package main

import (
	"github.com/gorilla/mux"
	"github.com/gowiki-api/Services"
	"github.com/gowiki-api/controllers"
	"net/http"
)

func main() {
	router := InitRouter()
	http.ListenAndServe(":8080", router)
}

func InitRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	// -------- Route Without Middleware Check --------//

	// router.Path("/").HandlerFunc(controllers.Main)
	//router.Methods("POST").Path("/user").Name("CreateUser").HandlerFunc(controllers.UserCreate)
	router.Methods("POST").Path("/auth").Name("auth").HandlerFunc(controllers.AuthUsers)

	var handler http.Handler
	AuthMiddleware := Services.AuthentificationMiddleware(handler)
	router.Use(AuthMiddleware)

	// -------- Route With Middleware Check --------//

	/*router.Methods("POST").Path("/article").Name("create").HandlerFunc(controllers.ArticleCreate)
	router.Methods("PUT").Path("/article/{id}").Name("Update").HandlerFunc(controllers.ArticleUpdate)
	router.Methods("POST").Path("/comment").Name("CreateComment").HandlerFunc(controllers.CommentCreate).Subrouter().Use(AuthMiddleware)
	*/

	router.Methods("POST").Path("/refresh").Name("refresh").HandlerFunc(Services.RefreshToken)
	/*
		router.Methods("GET").Path("/article/{id}/view").Name("View").HandlerFunc(controllers.ArticleView)
		router.Methods("GET").Path("/user/{id}").Name("ViewUser").HandlerFunc(controllers.UserView)*/

	return router
}
