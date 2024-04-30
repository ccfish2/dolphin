package cell

import (
	"context"
)

// HookContext is a context passed to a lifecycle hook that is
// cancelled in case of timeout
type HookContext context.Context
