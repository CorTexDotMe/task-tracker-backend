package server

import (
	"log"
	"net/http"
	"os"

	"task-tracker-backend/internal/database"
	"task-tracker-backend/internal/graph"
	"task-tracker-backend/internal/middleware"
	"task-tracker-backend/internal/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
)

// TODO env for main port
const defaultPort = "8080"

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	database.InitDB()
	database.Migrate()

	router := chi.NewRouter()
	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.Recoverer)
	router.Use(chiMiddleware.Logger)

	router.Group(func(r chi.Router) {
		r.Handle("/", playground.Handler("GraphQL playground", "/query"))
		r.Post("/login", resolver.Login)
		r.Post("/register", resolver.Register)
	})

	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))
	router.Group(func(r chi.Router) {
		r.Use(middleware.OnlyAuthenticated())
		r.Handle("/query", server)
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
