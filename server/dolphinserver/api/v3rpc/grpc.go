package v3rpc

import (
	"crypto/tls"

	"github.com/ccfish2/dolphin/server/dolphinserver"
	"google.golang.org/grpc"
)

func Server(s *dolphinserver.DolphinServer, tls *tls.Config, interceptor grpc.UnaryServerInterceptor, opts ...grpc.ServerOption) *grpc.Server {
	return nil
}
