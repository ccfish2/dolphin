package api

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

type AdminDisableHandler struct {
	name string
}

func NewAdminDisableHandler(name string) *AdminDisableHandler {
	return &AdminDisableHandler{name: name}
}

func (a *AdminDisableHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	log.Infof("API %s is disabled", a.name)
	_, _ = w.Write([]byte("This API is disabled"))
}

func DisableAPIs(paths PathSet, addMiddleWare func(method, path string, builder middleware.Builder)) {
	for k, pm := range paths {
		addMiddleWare(pm.Method, pm.Path, func(_ http.Handler) http.Handler {
			return NewAdminDisableHandler(k)
		})
	}
}
