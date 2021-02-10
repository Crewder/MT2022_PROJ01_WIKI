package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gowiki-api/config"
	"github.com/gowiki-api/controllers"
)

func main() {
	config.DatabaseInit()
	router := InitRouter()
	http.ListenAndServe(":8080", router)
}

func InitRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.Path("/").HandlerFunc(controllers.Main)

	/*router.Methods("POST").Path("/article").Name("create").HandlerFunc(controllers.ArticleCreate)
	router.Methods("PUT").Path("/article/{id}").Name("Update").HandlerFunc(controllers.ArticleUpdate)
	router.Methods("POST").Path("/comment").Name("CreateComment").HandlerFunc(controllers.CommentCreate)
	router.Methods("POST").Path("/user").Name("CreateUser").HandlerFunc(controllers.UserCreate)
	router.Methods("POST").Path("/auth").Name("Auth").HandlerFunc(controllers.UserAuth)
	router.Methods("GET").Path("/article/{id}/view").Name("View").HandlerFunc(controllers.ArticleView)
	router.Methods("GET").Path("/user/{id}").Name("ViewUser").HandlerFunc(controllers.UserView)*/

	return router

}
