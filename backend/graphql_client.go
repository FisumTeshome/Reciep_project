package main

// import (
// 	"net/http"

// 	"github.com/shurcooL/graphql"
// )

// // GraphQLClient is the global GraphQL client instance
// var GraphQLClient *graphql.Client

// // InitializeGraphQLClient sets up the GraphQL client
// func InitializeGraphQLClient() {
// 	// Define Hasura GraphQL API endpoint
// 	hasuraEndpoint := "http://localhost:8080/v1/graphql"

// 	// Create HTTP client with Hasura admin secret (if enabled)
// 	httpClient := &http.Client{
// 		Transport: &authTransport{
// 			Transport: http.DefaultTransport,
// 			Secret:    "myadminsecretkey", // Use env variable for production
// 		},
// 	}

// 	// Initialize GraphQL client
// 	GraphQLClient = graphql.NewClient(hasuraEndpoint, httpClient)
// }

// // authTransport is a custom HTTP transport that adds authentication headers
// type authTransport struct {
// 	Transport http.RoundTripper
// 	Secret    string
// }

// func (at *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("X-Hasura-Admin-Secret", at.Secret) // Use JWT in production
// 	return at.Transport.RoundTrip(req)
// }
