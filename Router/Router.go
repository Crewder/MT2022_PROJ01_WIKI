package Router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/gowiki-api/Middleware"
	"github.com/gowiki-api/Services"
	"github.com/gowiki-api/controllers"
	"net/http"
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
	router.Post("/login", controllers.AuthUsers)       //		"/auth" - Authentificate by credentials
	router.Post("/createuser", controllers.CreateUser) // 		"/Createuser" - Create a User and hash password
	router.Post("/logout", controllers.Logout)         // 		"/Logout" - Disconnect and suppress token
	//router.Get("/article/{id}/view", controllers.ArticleView)
	//router.Post.("/user",controllers.UserCreate)

	// -------- Private Route  --------//
	PrivateRouter := router.Group(nil)
	PrivateRouter.Use(Middleware.AuthentificationMiddleware) // Verify the JwtToken
	PrivateRouter.Use(Middleware.CSRFMiddleware)             // Verify The CSRF TOKEN

	PrivateRouter.Post("/refresh", Services.RefreshToken) // 	"/refresh" - refresh the Token
	// PrivateRouter.Post.("/article", controllers.ArticleCreate)
	// PrivateRouter.Put.("/article/{id}", controllers.ArticleUpdate)
	// PrivateRouter.Post.("/comment", controllers.CommentCreate)
	// PrivateRouter.Get("/article/{id}/view", controllers.ArticleView)
	// PrivateRouter.Get("/user/{id}",controllers.UserView)

	return router
}
