package main

import (
	"github.com/gorilla/mux"
	"github.com/gowiki-api/controllers"
	"net/http"
)

func main() {
	router := InitRouter()
	http.ListenAndServe(":8080", router)
}

func InitRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/article/{id}/view").Name("View Article").HandlerFunc(controllers.GetArticle)
	router.Methods("GET").Path("/articles/view").Name("View Articles").HandlerFunc(controllers.GetArticles)
	router.Methods("POST").Path("/comment/create").Name("CreateComment").HandlerFunc(controllers.CommentCreate)
	router.Methods("POST").Path("/article").Name("create").HandlerFunc(controllers.ArticleCreate)

	router.Methods("GET").Path("/article/{id}/view").Name("View Article").HandlerFunc(controllers.GetArticle)
	router.Methods("GET").Path("/articles/view").Name("View Articles").HandlerFunc(controllers.GetArticles)

	/*router.Methods("POST").Path("/article").Name("create").HandlerFunc(controllers.ArticleCreate)
	router.Methods("PUT").Path("/article/{id}").Name("Update").HandlerFunc(controllers.ArticleUpdate)
	router.Methods("POST").Path("/comment").Name("CreateComment").HandlerFunc(controllers.CommentCreate)
	router.Methods("POST").Path("/user").Name("CreateUser").HandlerFunc(controllers.UserCreate)
	router.Methods("POST").Path("/auth").Name("Auth").HandlerFunc(controllers.UserAuth)
	router.Methods("GET").Path("/article/{id}/view").Name("View").HandlerFunc(controllers.ArticleView)
	router.Methods("GET").Path("/user/{id}").Name("ViewUser").HandlerFunc(controllers.UserView)*/

	return router

}
