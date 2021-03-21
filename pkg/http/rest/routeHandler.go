package rest

import (
	"github.com/go-chi/chi"
	"github.com/gowiki-api/pkg/auth/jwt"
	"github.com/gowiki-api/pkg/handler"
	"github.com/gowiki-api/pkg/http/middleware"
	_ "github.com/gowiki-api/pkg/http/middleware"
	"net/http"
)

func Router() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.CORSMiddleware)

	// -------- Anonymous route  --------//
	router.Post("/user/login", jwt.AuthUsers)
	router.Post("/user/create", handler.CreateUser)
	router.Get("/article/{slug}", handler.GetArticle)
	router.Get("/articles", handler.GetArticles)

	// -------- Private Route  --------//
	// -------- Config
	PrivateRouter := router.Group(nil)

	PrivateRouter.Use(middleware.AuthentificationMiddleware)

	// -------- Private Route
	PrivateRouter.Post("/article/create", handler.CreateArticle)
	PrivateRouter.Put("/article/{slug}", handler.UpdateArticle)
	PrivateRouter.Post("/comment/create", handler.CreateComment)
	PrivateRouter.Post("/user/logout", jwt.Logout)
	PrivateRouter.Get("/comment/{id}", handler.GetCommentsByArticle)

	// -------- Admin Route
	PrivateRouter.Get("/users", handler.GetUsers)
	PrivateRouter.Get("/user/{id}", handler.GetUser)
	//PrivateRouter.Put("/role", handler.UpdateRole)

	return router
}
