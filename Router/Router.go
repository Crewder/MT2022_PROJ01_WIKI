package Router

import (
	"github.com/go-chi/chi"
	"github.com/gowiki-api/Middleware"
	"github.com/gowiki-api/Services"
	"github.com/gowiki-api/controllers"
	"net/http"
)

func Router() http.Handler {
	router := chi.NewRouter()

	// -------- Public route  --------//

	router.Post("/auth", controllers.AuthUsers)        //		"/auth" - Authentificate by credentials
	router.Post("/createuser", controllers.CreateUser) // " /Createuser" - Create a User and hash password

	//router.GET("/article/{id}/view", controllers.ArticleView)
	//router.POST.("/user",controllers.UserCreate)

	// -------- Private Route  --------//
	// Todo middleware Cors

	PrivateRouter := router.Group(nil)
	PrivateRouter.Use(Middleware.AuthentificationMiddleware) // Verify the JwtToken
	PrivateRouter.Use(Middleware.CSRFMiddleware)             // Verify The CSRF TOKEN

	PrivateRouter.Post("/refresh", Services.RefreshToken) //	// 	"/refresh" - refresh the Token

	// PrivateRouter.POST.("/article", controllers.ArticleCreate)
	// PrivateRouter.PUT.("/article/{id}", controllers.ArticleUpdate)
	// PrivateRouter.POST.("/comment", controllers.CommentCreate)
	// PrivateRouter.GET("/article/{id}/view", controllers.ArticleView)
	// PrivateRouter.GET("/user/{id}",controllers.UserView)

	return router
}
