package services

import (
	"errors"
	"strings"

	"github.com/FisumTeshome/Reciep_project/database"
	"github.com/FisumTeshome/Reciep_project/models"
	"github.com/google/uuid"
)

type IngredientInput struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
	Notes    string  `json:"notes"`
}

type StepInput struct {
	StepNumber   int    `json:"step_number"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	TimeRequired int    `json:"time_required"`
}

type CreateRecipeInput struct {
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	CategoryID    string          `json:"category_id"`
	PrepTime      int             `json:"prep_time"`
	CookTime      int             `json:"cook_time"`
	Servings      int             `json:"servings"`
	Difficulty    string          `json:"difficulty"`
	FeaturedImage string          `json:"featured_image"`
	Images        []string        `json:"images"`
	Ingredients   []IngredientInput `json:"ingredients"`
	Steps         []StepInput     `json:"steps"`
	IsPublished   bool            `json:"is_published"`
}

func ListRecipes(search string, categoryID string) ([]models.Recipe, error) {
	var recipes []models.Recipe
	query := database.DB.Model(&models.Recipe{}).
		Preload("User").
		Preload("Category").
		Preload("Steps").
		Preload("RecipeIngredients").
		Preload("RecipeIngredients.Ingredient").
		Order("created_at DESC")

	if strings.TrimSpace(search) != "" {
		likeTerm := "%" + search + "%"
		query = query.Where("title ILIKE ? OR description ILIKE ?", likeTerm, likeTerm)
	}
	if strings.TrimSpace(categoryID) != "" {
		parsed, err := uuid.Parse(categoryID)
		if err != nil {
			return nil, err
		}
		query = query.Where("category_id = ?", parsed)
	}

	if err := query.Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}

func GetRecipeByID(recipeID string) (models.Recipe, error) {
	parsed, err := uuid.Parse(recipeID)
	if err != nil {
		return models.Recipe{}, err
	}

	var recipe models.Recipe
	if err := database.DB.Preload("User").
		Preload("Category").
		Preload("Steps").
		Preload("RecipeIngredients").
		Preload("RecipeIngredients.Ingredient").
		First(&recipe, "id = ?", parsed).Error; err != nil {
		return models.Recipe{}, err
	}
	return recipe, nil
}

func CreateRecipe(userID string, input CreateRecipeInput) (models.Recipe, error) {
	if strings.TrimSpace(input.Title) == "" {
		return models.Recipe{}, errors.New("recipe title is required")
	}
	if strings.TrimSpace(input.FeaturedImage) == "" {
		return models.Recipe{}, errors.New("featured image is required")
	}
	if strings.TrimSpace(input.CategoryID) == "" {
		return models.Recipe{}, errors.New("category is required")
	}

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return models.Recipe{}, err
	}

	parsedCategoryID, err := uuid.Parse(input.CategoryID)
	if err != nil {
		return models.Recipe{}, err
	}

	var category models.Category
	if err := database.DB.First(&category, "id = ?", parsedCategoryID).Error; err != nil {
		return models.Recipe{}, err
	}

	recipe := models.Recipe{
		UserID:        parsedUserID,
		CategoryID:    parsedCategoryID,
		Title:         strings.TrimSpace(input.Title),
		Description:   input.Description,
		PrepTime:      input.PrepTime,
		CookTime:      input.CookTime,
		Servings:      input.Servings,
		Difficulty:    input.Difficulty,
		FeaturedImage: input.FeaturedImage,
		Images:        input.Images,
		IsPublished:   input.IsPublished,
	}

	if recipe.Difficulty == "" {
		recipe.Difficulty = "medium"
	}
	if recipe.Servings <= 0 {
		recipe.Servings = 2
	}

	if err := database.DB.Create(&recipe).Error; err != nil {
		return models.Recipe{}, err
	}

	for _, ingredientInput := range input.Ingredients {
		ingredientName := strings.TrimSpace(ingredientInput.Name)
		if ingredientName == "" {
			continue
		}

		var ingredient models.Ingredient
		if err := database.DB.Where("name ILIKE ?", ingredientName).FirstOrCreate(&ingredient, models.Ingredient{Name: ingredientName}).Error; err != nil {
			return models.Recipe{}, err
		}

		recipeIngredient := models.RecipeIngredient{
			RecipeID:     recipe.ID,
			IngredientID: ingredient.ID,
			Quantity:     ingredientInput.Quantity,
			Unit:         ingredientInput.Unit,
			Notes:        ingredientInput.Notes,
		}

		if err := database.DB.Create(&recipeIngredient).Error; err != nil {
			return models.Recipe{}, err
		}
	}

	for _, stepInput := range input.Steps {
		stepDescription := strings.TrimSpace(stepInput.Description)
		if stepDescription == "" {
			continue
		}

		step := models.RecipeStep{
			RecipeID:     recipe.ID,
			StepNumber:   stepInput.StepNumber,
			Description:  stepDescription,
			Image:        stepInput.Image,
			TimeRequired: stepInput.TimeRequired,
		}
		if step.StepNumber <= 0 {
			step.StepNumber = 1
		}
		if err := database.DB.Create(&step).Error; err != nil {
			return models.Recipe{}, err
		}
	}

	return GetRecipeByID(recipe.ID.String())
}

func UpdateRecipe(userID string, recipeID string, input CreateRecipeInput) (models.Recipe, error) {
	parsedRecipeID, err := uuid.Parse(recipeID)
	if err != nil {
		return models.Recipe{}, err
	}

	var recipe models.Recipe
	if err := database.DB.First(&recipe, "id = ?", parsedRecipeID).Error; err != nil {
		return models.Recipe{}, err
	}

	if recipe.UserID.String() != userID {
		return models.Recipe{}, errors.New("you can only update your own recipes")
	}

	recipe.Title = input.Title
	recipe.Description = input.Description
	recipe.PrepTime = input.PrepTime
	recipe.CookTime = input.CookTime
	recipe.Servings = input.Servings
	recipe.Difficulty = input.Difficulty
	recipe.FeaturedImage = input.FeaturedImage
	recipe.Images = input.Images
	recipe.IsPublished = input.IsPublished

	if strings.TrimSpace(recipe.Difficulty) == "" {
		recipe.Difficulty = "medium"
	}
	if recipe.Servings <= 0 {
		recipe.Servings = 2
	}

	if err := database.DB.Save(&recipe).Error; err != nil {
		return models.Recipe{}, err
	}

	if err := database.DB.Where("recipe_id = ?", recipe.ID).Delete(&models.RecipeIngredient{}).Error; err != nil {
		return models.Recipe{}, err
	}
	if err := database.DB.Where("recipe_id = ?", recipe.ID).Delete(&models.RecipeStep{}).Error; err != nil {
		return models.Recipe{}, err
	}

	for _, ingredientInput := range input.Ingredients {
		ingredientName := strings.TrimSpace(ingredientInput.Name)
		if ingredientName == "" {
			continue
		}
		var ingredient models.Ingredient
		if err := database.DB.Where("name ILIKE ?", ingredientName).FirstOrCreate(&ingredient, models.Ingredient{Name: ingredientName}).Error; err != nil {
			return models.Recipe{}, err
		}

		if err := database.DB.Create(&models.RecipeIngredient{
			RecipeID:     recipe.ID,
			IngredientID: ingredient.ID,
			Quantity:     ingredientInput.Quantity,
			Unit:         ingredientInput.Unit,
			Notes:        ingredientInput.Notes,
		}).Error; err != nil {
			return models.Recipe{}, err
		}
	}

	for _, stepInput := range input.Steps {
		stepDescription := strings.TrimSpace(stepInput.Description)
		if stepDescription == "" {
			continue
		}
		if err := database.DB.Create(&models.RecipeStep{
			RecipeID:     recipe.ID,
			StepNumber:   stepInput.StepNumber,
			Description:  stepDescription,
			Image:        stepInput.Image,
			TimeRequired: stepInput.TimeRequired,
		}).Error; err != nil {
			return models.Recipe{}, err
		}
	}

	return GetRecipeByID(recipe.ID.String())
}

func DeleteRecipe(userID string, recipeID string) error {
	parsedRecipeID, err := uuid.Parse(recipeID)
	if err != nil {
		return err
	}

	var recipe models.Recipe
	if err := database.DB.First(&recipe, "id = ?", parsedRecipeID).Error; err != nil {
		return err
	}
	if recipe.UserID.String() != userID {
		return errors.New("you can only delete your own recipes")
	}

	return database.DB.Where("id = ?", recipe.ID).Delete(&models.Recipe{}).Error
}

func GetRecipesByUser(userID string) ([]models.Recipe, error) {
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	var recipes []models.Recipe
	if err := database.DB.Where("user_id = ?", parsedUserID).
		Preload("User").
		Preload("Category").
		Preload("Steps").
		Preload("RecipeIngredients").
		Preload("RecipeIngredients.Ingredient").
		Order("created_at DESC").
		Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}

func GetCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := database.DB.Order("name ASC").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoriesByName(name string) ([]models.Category, error) {
	var categories []models.Category
	if err := database.DB.Where("name ILIKE ?", "%"+name+"%").Order("name ASC").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}