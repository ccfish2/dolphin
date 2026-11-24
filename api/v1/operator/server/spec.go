package server

import (
	"github.com/go-openapi/loads"

	api "github.com/ccfish2/dolphin/pkg/api"
)

type Spec struct {
	*loads.Document

	// A set of APIs that are disabled.
	DeniedAPIs api.PathSet
}

func newSpec(allowedAPIs []string) (*Spec, error) {
	swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
	if err != nil {
		return nil, err
	}

	deniedAPI, err := api.AllowedFlagsToDeniedPaths(
		swaggerSpec,
		allowedAPIs,
	)
	if err != nil {
		return nil, err
	}

	return &Spec{
		Document:   swaggerSpec,
		DeniedAPIs: deniedAPI,
	}, nil
}
