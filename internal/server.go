package server

import (
	"log"
	"net/http"
	"os"

	"task-tracker-backend/internal/graph"
	"task-tracker-backend/internal/middleware"
	"task-tracker-backend/internal/resolver"
	"task-tracker-backend/internal/utils"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
)

// Run the server
//
// Load dotenv properties, initialize database, define router
// with routes and middleware, and start the server
func Run() {
	err := godotenv.Load("../../.env")
	utils.HandleError(err)
	port := os.Getenv("SERVER_PORT")

	router := chi.NewRouter()
	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.Recoverer)
	router.Use(chiMiddleware.Logger)

	router.Group(func(r chi.Router) {
		r.Handle("/", playground.Handler("GraphQL playground", "/query"))
		r.Post("/login", resolver.Login)
		r.Post("/register", resolver.Register)
	})

	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver.NewResolver()}))
	router.Group(func(r chi.Router) {
		r.Use(middleware.OnlyAuthenticated())
		r.Handle("/query", server)
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
