package rest

import (
	"github.com/go-chi/chi"
	"github.com/gowiki-api/wiki/auth/jwt"
	"github.com/gowiki-api/wiki/controllers"
	"github.com/gowiki-api/wiki/http/middleware"
	_ "github.com/gowiki-api/wiki/http/middleware"
	"net/http"
)

func Router() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.CORSMiddleware)

	// -------- Anonymous route  --------//
	router.Post("/user/login", jwt.AuthUsers)
	router.Post("/user/create", controllers.CreateUser)
	router.Get("/article/{slug}", controllers.GetArticle)
	router.Get("/articles", controllers.GetArticles)

	// -------- Private Route  --------//
	// -------- Config
	PrivateRouter := router.Group(nil)
	PrivateRouter.Use(middleware.AuthentificationMiddleware)

	// -------- Private Route
	PrivateRouter.Post("/article/create", controllers.CreateArticle)
	PrivateRouter.Put("/article/{slug}", controllers.UpdateArticle)
	PrivateRouter.Delete("/article/{slug}", controllers.DeleteArticle)
	PrivateRouter.Post("/comment/create", controllers.CreateComment)
	PrivateRouter.Post("/user/logout", jwt.Logout)
	PrivateRouter.Get("/comment/{id}", controllers.GetCommentsByArticle)
	PrivateRouter.Put("/comment/{id}", controllers.UpdateComment)
	PrivateRouter.Delete("/comment/{id}", controllers.DeleteComment)

	// -------- Admin Route
	PrivateRouter.Put("/comment/{id}", controllers.UpdateComment)
	PrivateRouter.Delete("/comment/{id}", controllers.DeleteComment)
	PrivateRouter.Get("/users", controllers.GetUsers)
	PrivateRouter.Get("/user/{id}", controllers.GetUser)

	return router
}
