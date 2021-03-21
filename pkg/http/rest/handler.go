package rest

import (
	"github.com/go-chi/chi"
	"github.com/gowiki-api/pkg/handler"
	"github.com/gowiki-api/pkg/http/middleware"
	"net/http"
)

func Router() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.CORSMiddleware) // Configure CORS

	// -------- Public route  --------//

	router.Post("/user/login", handler.AuthUsers)
	router.Post("/user/create", handler.CreateUser)
	router.Post("/user/logout", handler.Logout)
	router.Get("/article/{slug}", handler.GetArticle)
	router.Get("/articles", handler.GetArticles)
	router.Get("/comment/{id}", handler.GetCommentsByArticle)

	// -------- Private Route  --------//
	PrivateRouter := router.Group(nil)
	PrivateRouter.Use(middleware.AuthentificationMiddleware) // Verify the JwtToken and CSRF

	PrivateRouter.Post("/article/create", handler.CreateArticle)
	PrivateRouter.Put("/article/{slug}", handler.UpdateArticle)
	PrivateRouter.Delete("/article/{slug}", handler.DeleteArticle)
	PrivateRouter.Post("/comment/create", handler.CreateComment)
	PrivateRouter.Delete("/comment/{id}", handler.DeleteComment)
	PrivateRouter.Get("/users", handler.GetUsers)
	PrivateRouter.Get("/user/{id}", handler.GetUser)

	return router
}
