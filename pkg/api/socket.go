package api

import (
	"os"
	"os/user"
	"strconv"

	"github.com/ccfish2/infra/pkg/logging"
	"github.com/ccfish2/infra/pkg/logging/logfields"
)

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, "api")

func getGrpIDByName(grpName string) (int, error) {

	grp, err := user.LookupGroup(grpName)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(grp.Gid)
}

func SetDefaultPermissions(socketPath string) error {
	gid, err := getGrpIDByName("dolphin")
	if err != nil {
		return err
	} else {
		if err := os.Chown(socketPath, 0, gid); err != nil {
			return err
		}
	}
	return nil
}
