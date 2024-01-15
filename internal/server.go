package server

import (
	"log"
	"net/http"
	"os"

	"task-tracker-backend/internal/database"
	"task-tracker-backend/internal/graph"
	"task-tracker-backend/internal/resolver"
	"task-tracker-backend/internal/security"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

// TODO env for main port
const defaultPort = "8080"

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(security.Filter())

	database.InitDB()
	// defer database.CloseDB()
	database.Migrate()

	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
