package main

import (
	"log"
	"net/http"

	"github.com/SnapshotApp/api/handler"
	"github.com/SnapshotApp/api/resolver"
	"github.com/SnapshotApp/api/schema"

	graphql "github.com/graph-gophers/graphql-go"
)

func main() {
	// Parse GraphQL schema
	s, err := schema.GetSchema("./schema/schema.graphql")
	if err != nil {
		log.Fatalf("Unable to parse GraphQL schema: %s\n", err)
	}
	schema := graphql.MustParseSchema(s, &resolver.Resolver{})

	// Setup GraphQL endpoint
	http.Handle("/graphql", &handler.Handler{Schema: schema})

	// Start the server
	port := ":8000"
	log.Printf("Server running on localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
