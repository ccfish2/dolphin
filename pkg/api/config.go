package api

import (
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

// Endpoint is an API Endpoint for a parsed API specification.
type Endpoint struct {
	Method      string
	Path        string
	Description string
}
type PathSet map[string]Endpoint

func parsesSpecPath(paths spec.Paths) PathSet {
	panic("")
}

func generateDeniedAPIEndpoints(allpaths PathSet, allowed []string) (PathSet, error) {
	panic("")
}

func AllowedFlagsToDeniedPaths(spec *loads.Document, allowed []string) (PathSet, error) {
	panic("")
}
