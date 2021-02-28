package rest

import (
	"github.com/go-chi/chi"
	"github.com/gowiki-api/pkg/handler"
	"github.com/gowiki-api/pkg/http/jwt"
	"github.com/gowiki-api/pkg/http/middleware"
	"net/http"
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.CORSMiddleware) // Configure CORS
	//router.Use(middleware.CSRFMiddleware)             // Verify The CSRF TOKEN

	// -------- Public route  --------//

	router.Post("/user/login", handler.AuthUsers)
	router.Post("/user/create", handler.AddUser)
	router.Post("/user/logout", handler.Logout)
	router.Get("/article/{id}", handler.GetArticle)
	router.Get("/articles", handler.GetArticles)

	// -------- Private Route  --------//
	PrivateRouter := router.Group(nil)
	PrivateRouter.Use(middleware.AuthentificationMiddleware) // Verify the JwtToken

	PrivateRouter.Post("/user/refresh", jwt.RefreshToken)
	PrivateRouter.Post("/article/create", handler.ArticleCreate)
	PrivateRouter.Put("/article/{id}", handler.ArticleUpdate)
	PrivateRouter.Post("/comment/create", handler.CommentCreate)
	PrivateRouter.Get("/users", handler.GetUsers)
	PrivateRouter.Get("/user/{id}", handler.GetUser)

	return router
}
