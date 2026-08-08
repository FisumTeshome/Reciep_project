package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/FisumTeshome/Reciep_project/services"
)

func ListRecipes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	search := r.URL.Query().Get("search")
	categoryID := r.URL.Query().Get("category_id")

	recipes, err := services.ListRecipes(search, categoryID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, recipes)
}

func GetRecipe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	recipeID := r.URL.Query().Get("id")
	if recipeID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "recipe id is required"})
		return
	}

	recipe, err := services.GetRecipeByID(recipeID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "recipe not found"})
		return
	}

	writeJSON(w, http.StatusOK, recipe)
}

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	tokenString, err := requireAuthToken(r)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	user, err := services.GetUserFromToken(tokenString)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
		return
	}

	var input services.CreateRecipeInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	recipe, err := services.CreateRecipe(user.ID.String(), input)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, recipe)
}

func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	tokenString, err := requireAuthToken(r)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	user, err := services.GetUserFromToken(tokenString)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
		return
	}

	recipeID := r.URL.Query().Get("id")
	if recipeID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "recipe id is required"})
		return
	}

	var input services.CreateRecipeInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	recipe, err := services.UpdateRecipe(user.ID.String(), recipeID, input)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, recipe)
}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	tokenString, err := requireAuthToken(r)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	user, err := services.GetUserFromToken(tokenString)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
		return
	}

	recipeID := r.URL.Query().Get("id")
	if recipeID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "recipe id is required"})
		return
	}

	if err := services.DeleteRecipe(user.ID.String(), recipeID); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "recipe deleted successfully"})
}