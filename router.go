package main

import (
	"github.com/gorilla/mux"
	"github.com/gowiki-api/controllers"
)

func InitRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.Methods("POST").Path("/article").Name("create").HandlerFunc(controllers.ArticleCreate)
	router.Methods("PUT").Path("/article/{id}").Name("Update").HandlerFunc(controllers.ArticleUpdate)
	router.Methods("POST").Path("/comment").Name("CreateComment").HandlerFunc(controllers.CommentCreate)
	router.Methods("POST").Path("/user").Name("CreateUser").HandlerFunc(controllers.UserCreate)
	router.Methods("POST").Path("/auth").Name("Auth").HandlerFunc(controllers.UserAuth)

	return router

}
