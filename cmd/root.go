package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"quesurifn/simple-auth/cmd/gql"
	"quesurifn/simple-auth/cmd/grpc"
)

var rootCmd = &cobra.Command{
	Use:   "simp",
	Short: "Simp is the backend for simple auth",
	Long: `Simple Auth is a fully managed R/ABAC service.
		Simple auth is designed to tackle the middle market between
			Open Source and expensive Auth0.

		We will get to 5m ARR within 3 years!
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	// Add subcommands to the root command
	rootCmd.AddCommand(grpc.GrpcCmd)
	rootCmd.AddCommand(gql.ApolloCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
