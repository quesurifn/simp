package grpc

import (
	"fmt"

	"github.com/spf13/cobra"
)

// GrpcCmd represents the grpc command
var GrpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "gRPC server commands",
	Long: `gRPC server commands allow you to work with gRPC API.
For example, you can generate gRPC stubs, run gRPC server, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grpc called")
	},
}
