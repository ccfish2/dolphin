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

func parsesSpecPaths(paths *spec.Paths) PathSet {
	res := make(PathSet)
	// impl if need
	return PathSet(res)
}

func generateDeniedAPIEndpoints(allpaths PathSet, allowed []string) (PathSet, error) {
	denied := allpaths
	// impl if need
	return denied, nil
}

func AllowedFlagsToDeniedPaths(spec *loads.Document, allowed []string) (PathSet, error) {
	paths := parsesSpecPaths(spec.Spec().Paths)
	return generateDeniedAPIEndpoints(paths, allowed)
}
