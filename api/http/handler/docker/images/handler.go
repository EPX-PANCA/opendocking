package images

import (
	"net/http"

	"github.com/opendocking/opendocking/api/docker/client"
	"github.com/opendocking/opendocking/api/http/middlewares"
	"github.com/opendocking/opendocking/api/http/security"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"

	"github.com/gorilla/mux"
)

type Handler struct {
	*mux.Router
	dockerClientFactory *client.ClientFactory
	bouncer             security.BouncerService
}

// NewHandler creates a handler to process non-proxied requests to docker APIs directly.
func NewHandler(routePrefix string, bouncer security.BouncerService, dockerClientFactory *client.ClientFactory) *Handler {
	h := &Handler{
		Router:              mux.NewRouter(),
		dockerClientFactory: dockerClientFactory,
		bouncer:             bouncer,
	}

	router := h.PathPrefix(routePrefix).Subrouter()
	router.Use(bouncer.AuthenticatedAccess, middlewares.CheckEndpointAuthorization(bouncer))

	router.Handle("", httperror.LoggerHandler(h.imagesList)).Methods(http.MethodGet)
	return h
}
