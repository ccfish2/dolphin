package dolphinadmin

import (
	"github.com/ccfish2/dolphin/server/embed"
)

func startDolphinOrV2(args []string) {

}

func startDolphin() (<-chan struct{}, <-chan struct{}, error) {
	_, err := embed.StartDophin(nil)
	return nil, nil, err
}
