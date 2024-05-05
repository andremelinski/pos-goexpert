//go:build tools
// +build tools

// https://gqlgen.com/
package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
