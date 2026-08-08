package routes

import (
"net/http"

"github.com/FisumTeshome/Reciep_project/controllers"
)

func registerRecipeRoutes(mux *http.ServeMux) {
mux.HandleFunc("/api/recipes", withCORS(controllers.ListRecipes))
mux.HandleFunc("/api/recipes/get", withCORS(controllers.GetRecipe))
mux.HandleFunc("/api/recipes/create", withCORS(controllers.CreateRecipe))
mux.HandleFunc("/api/recipes/update", withCORS(controllers.UpdateRecipe))
mux.HandleFunc("/api/recipes/delete", withCORS(controllers.DeleteRecipe))
}
