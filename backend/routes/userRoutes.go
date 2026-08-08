package routes

import (
	"net/http"

	"github.com/FisumTeshome/Reciep_project/controllers"
)

func registerUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/users/me", withCORS(controllers.GetProfile))
	mux.HandleFunc("/api/users/profile", withCORS(controllers.UpdateProfile))
	mux.HandleFunc("/api/users/password", withCORS(controllers.UpdatePassword))
}