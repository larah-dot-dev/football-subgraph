package main

import (
	"football-subgraph/graph"
	"os"

	"football-subgraph/football_data"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

func background_fetch_match_data() {
	for {
		football_data.GetMatches()
		time.Sleep(time.Hour)
	}
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	go background_fetch_match_data()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	gqlgen_server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	playground_server := playground.Handler("GraphQL", "/query")

	router.POST("/query", func(c *gin.Context) {
		gqlgen_server.ServeHTTP(c.Writer, c.Request)
	})

	router.GET("/", func(c *gin.Context) {
		playground_server.ServeHTTP(c.Writer, c.Request)
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	router.SetTrustedProxies(nil)
	router.Run(":" + port)
}
