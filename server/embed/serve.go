package embed

import (
	"github.com/ccfish2/dolphin/server/dolphinserver"
	"github.com/ccfish2/dolphin/server/dolphinserver/api/v3rpc"

	"google.golang.org/grpc"
)

type serveCtx struct {
}

func newServeCtx() *serveCtx {
	return nil
}

func (sctx *serveCtx) serve(
	srv *dolphinserver.DolphinServer,
	gopts grpc.ServerOption,
) {
	// bump server supporting rpc protocol
	// it could serve server and client rquest
	// consumed by StartDolphin API call
	v3rpc.Server(srv, nil, nil, gopts)
}
