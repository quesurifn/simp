package gql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/cobra"
	"quesurifn/simple-auth/graph"
)

// apolloCmd represents the apollo command
var ApolloCmd = &cobra.Command{
	Use:   "apollo",
	Short: "Apollo GraphQL server commands",
	Long: `Apollo GraphQL server commands allow you to work with GraphQL API.
For example, you can generate GraphQL schema, run GraphQL server, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := fiber.New()
		h := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
		app.Post("/graphql", func(c *fiber.Ctx) error {
			wrapHandler(h.ServeHTTP)(c)
			return nil
		})

		app.Get("/playground", func(c *fiber.Ctx) error {
			wrapHandler(playground.Handler("GraphQL", "/graphql"))(c)
			return nil
		})

		// Start the server
		err := app.Listen(":3000")
		if err != nil {
			return
		}
	},
}
