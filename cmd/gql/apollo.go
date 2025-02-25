package gql

import (
	"fmt"

	"github.com/spf13/cobra"
)

// apolloCmd represents the apollo command
var ApolloCmd = &cobra.Command{
	Use:   "apollo",
	Short: "Apollo GraphQL server commands",
	Long: `Apollo GraphQL server commands allow you to work with GraphQL API.
For example, you can generate GraphQL schema, run GraphQL server, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apollo called")
	},
}
