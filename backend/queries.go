package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/shurcooL/graphql"
// )

// // Recipe represents the GraphQL recipe model
// type Recipe struct {
// 	ID          graphql.String
// 	Title       graphql.String
// 	Description graphql.String
// 	Category    graphql.String
// 	PrepTime    graphql.Int
// }

// // FetchRecipes fetches recipes from Hasura GraphQL API
// func FetchRecipes() {
// 	var query struct {
// 		Recipes []Recipe `graphql:"recipes { id title description category prepTime }"`
// 	}

// 	// Execute the query
// 	err := GraphQLClient.Query(context.Background(), &query, nil)
// 	if err != nil {
// 		log.Fatalf("Error fetching recipes: %v", err)
// 	}

// 	// Print fetched recipes
// 	for _, recipe := range query.Recipes {
// 		fmt.Printf("Recipe: %s - %s\n", recipe.Title, recipe.Description)
// 	}
// }
