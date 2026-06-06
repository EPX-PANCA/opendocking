package webhooks

import (
	"net/http"

	"github.com/opendocking/opendocking/api/dataservices"
	dockerclient "github.com/opendocking/opendocking/api/docker/client"
	"github.com/opendocking/opendocking/api/http/security"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"

	"github.com/gorilla/mux"
)

// Handler is the HTTP handler used to handle webhook operations.
type Handler struct {
	*mux.Router
	requestBouncer      security.BouncerService
	DataStore           dataservices.DataStore
	DockerClientFactory *dockerclient.ClientFactory
}

// NewHandler creates a handler to manage webhooks operations.
func NewHandler(bouncer security.BouncerService) *Handler {
	h := &Handler{
		Router:         mux.NewRouter(),
		requestBouncer: bouncer,
	}
	h.Handle("/webhooks",
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.webhookCreate))).Methods(http.MethodPost)
	h.Handle("/webhooks/{id}",
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.webhookUpdate))).Methods(http.MethodPut)
	h.Handle("/webhooks",
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.webhookList))).Methods(http.MethodGet)
	h.Handle("/webhooks/{id}",
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.webhookDelete))).Methods(http.MethodDelete)
	h.Handle("/webhooks/{token}",
		bouncer.PublicAccess(httperror.LoggerHandler(h.webhookExecute))).Methods(http.MethodPost)

	return h
}
