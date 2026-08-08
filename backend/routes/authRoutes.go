package routes

import (
	"net/http"

	"github.com/FisumTeshome/Reciep_project/controllers"
)

func registerAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/auth/register", withCORS(controllers.Register))
	mux.HandleFunc("/api/auth/login", withCORS(controllers.Login))
	mux.HandleFunc("/api/auth/logout", withCORS(controllers.Logout))
	mux.HandleFunc("/api/auth/me", withCORS(controllers.Me))
}