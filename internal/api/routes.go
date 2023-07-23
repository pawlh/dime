package api

import (
	"dime/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	// GraphQL
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	// TODO: only allow playground in dev mode
	e.GET("/playground", echo.WrapHandler(playground.Handler("GraphQL playground", "/graphql")))
	e.POST("/graphql", echo.WrapHandler(srv))

	// Default
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

}
