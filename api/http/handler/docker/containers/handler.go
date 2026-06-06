package containers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/opendocking/opendocking/api/dataservices"
	"github.com/opendocking/opendocking/api/docker"
	dockerclient "github.com/opendocking/opendocking/api/docker/client"
	"github.com/opendocking/opendocking/api/http/middlewares"
	"github.com/opendocking/opendocking/api/http/security"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"
)

type Handler struct {
	*mux.Router
	dockerClientFactory *dockerclient.ClientFactory
	dataStore           dataservices.DataStore
	containerService    *docker.ContainerService
	bouncer             security.BouncerService
}

// NewHandler creates a handler to process non-proxied requests to docker APIs directly.
func NewHandler(routePrefix string, bouncer security.BouncerService, dataStore dataservices.DataStore, dockerClientFactory *dockerclient.ClientFactory, containerService *docker.ContainerService) *Handler {
	h := &Handler{
		Router:              mux.NewRouter(),
		dataStore:           dataStore,
		dockerClientFactory: dockerClientFactory,
		containerService:    containerService,
		bouncer:             bouncer,
	}

	router := h.PathPrefix(routePrefix).Subrouter()
	router.Use(bouncer.AuthenticatedAccess, middlewares.CheckEndpointAuthorization(bouncer))

	router.Handle("/{containerId}/gpus", httperror.LoggerHandler(h.containerGpusInspect)).Methods(http.MethodGet)
	router.Handle("/{containerId}/recreate", httperror.LoggerHandler(h.recreate)).Methods(http.MethodPost)

	return h
}
