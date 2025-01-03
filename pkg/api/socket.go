package api

import (
	"github.com/ccfish2/infra/pkg/logging"
	"github.com/ccfish2/infra/pkg/logging/logfields"
)

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, "api")

func SetDefaultPermissions(socketPath string) error {
	return nil
}
