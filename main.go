package main

import (
	"github.com/go-chi/chi"
	"github.com/gowiki-api/Middleware"
	"github.com/gowiki-api/Services"
	"github.com/gowiki-api/controllers"
	"net/http"
)

func main() {
	router := Router()
	http.ListenAndServe(":8080", router)
}

func Router() http.Handler {

	router := chi.NewRouter()

	// -------- Public route  --------//

	router.Post("/auth", controllers.AuthUsers) //"/auth" - Authentificate by credentials

	// -------- Private Route  --------//
	// Todo middleware Cors - XSS - CSRF

	PrivateRouter := router.Group(nil)
	PrivateRouter.Use(Middleware.AuthentificationMiddleware)

	PrivateRouter.Post("/refresh", Services.RefreshToken) // "/refresh" - refresh the Token

	return router
}
