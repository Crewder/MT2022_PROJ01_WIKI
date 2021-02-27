package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/gowiki-api/controllers"
	"github.com/gowiki-api/middleware"
	"github.com/gowiki-api/services"
)

func Router() http.Handler {
	router := chi.NewRouter()

	// Configuration du Cross-origin resource sharing
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	// -------- Public route  --------//
	router.Post("/user/login", controllers.AuthUsers)   //		"/auth" - Authentificate by credentials
	router.Post("/user/create", controllers.CreateUser) // 		"/Createuser" - Create a User and hash password
	router.Post("/user/logout", controllers.Logout)     // 		"/Logout" - Disconnect and suppress token
	router.Get("/article/{id}", controllers.GetArticle)
	router.Get("/articles", controllers.GetArticles)

	// -------- Private Route  --------//
	PrivateRouter := router.Group(nil)
	PrivateRouter.Use(middleware.AuthentificationMiddleware) // Verify the JwtToken
	PrivateRouter.Use(middleware.CSRFMiddleware)             // Verify The CSRF TOKEN

	router.Post("/user/refresh", services.RefreshToken) // 	"/refresh" - refresh the Token
	router.Post("/article/create", controllers.ArticleCreate)
	router.Put("/article/{id}", controllers.ArticleUpdate)
	router.Post("/comment/create", controllers.CommentCreate)
	router.Get("/users", controllers.GetUsers)
	router.Get("/user/{id}", controllers.GetUser)

	return router
}
