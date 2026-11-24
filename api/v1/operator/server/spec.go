package server

import (
	"github.com/go-openapi/loads"
	"github.com/spf13/pflag"

	api "github.com/ccfish2/dolphin/pkg/api"
	"github.com/ccfish2/infra/pkg/hive/cell"
)

type ServerConfig struct {
	EnableDolphinOperatorServerAccess []string
}

var (
	defaultServerConfig = ServerConfig{
		[]string{"*"},
	}
	AdminEnableFlag = "enable-dolphin-operator-server-access"
)

func (def ServerConfig) Flags(flags *pflag.FlagSet) {
	flags.StringSlice(AdminEnableFlag, def.EnableDolphinOperatorServerAccess, "lists of allowed server access")
}

var SpecCell = cell.Module(
	"dolphin-operator-spec",
	"Dolphin Operator Specification",
	cell.Config(defaultServerConfig),
	cell.Provide(newSpec),
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
