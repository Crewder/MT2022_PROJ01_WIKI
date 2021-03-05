package rest

import (
	"github.com/casbin/casbin"
	"github.com/casbin/chi-authz"
	"github.com/go-chi/chi"
	"github.com/gowiki-api/pkg/auth/jwt"
	"github.com/gowiki-api/pkg/handler"
	"github.com/gowiki-api/pkg/http/middleware"
	"net/http"
)

func Router() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.CORSMiddleware) // Configure CORS
	e := casbin.NewEnforcer("pkg/auth/roles/auth_model.conf", "pkg/auth/roles/auth_policy.csv")
	router.Use(authz.Authorizer(e))

	// -------- Anonymous route  --------//
	router.Post("/user/login", jwt.AuthUsers)
	router.Post("/user/create", handler.CreateUser)
	router.Get("/article/{slug}", handler.GetArticle)
	router.Get("/articles", handler.GetArticles)

	// -------- Public route  --------//
	router.Post("/user/logout", jwt.Logout)
	router.Get("/comment/{id}", handler.GetCommentsByArticle)

	// -------- Private Route  --------//
	// -------- Config
	PrivateRouter := router.Group(nil)
	PrivateRouter.Use(middleware.TokenAuthenMiddleware) // Verify the JwtToken and CSRF

	// -------- Private Route
	PrivateRouter.Post("/article/create", handler.CreateArticle)
	PrivateRouter.Put("/article/{slug}", handler.UpdateArticle)
	PrivateRouter.Post("/comment/create", handler.CreateComment)

	// -------- Admin Route
	PrivateRouter.Get("/users", handler.GetUsers)
	PrivateRouter.Get("/user/{id}", handler.GetUser)

	return router
}
