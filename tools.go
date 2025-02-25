//go:build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	_ "github.com/martinxsliu/protoc-gen-graphql"
	_ "github.com/spf13/cobra"
)
