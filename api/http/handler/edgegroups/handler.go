package edgegroups

import (
	"net/http"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	"github.com/opendocking/opendocking/api/http/security"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"

	"github.com/gorilla/mux"
)

// Handler is the HTTP handler used to handle environment(endpoint) group operations.
type Handler struct {
	*mux.Router
	DataStore            dataservices.DataStore
	ReverseTunnelService portainer.ReverseTunnelService
}

// NewHandler creates a handler to manage environment(endpoint) group operations.
func NewHandler(bouncer security.BouncerService) *Handler {
	h := &Handler{
		Router: mux.NewRouter(),
	}
	h.Handle("/edge_groups",
		bouncer.AdminAccess(bouncer.EdgeComputeOperation(httperror.LoggerHandler(h.edgeGroupCreate)))).Methods(http.MethodPost)
	h.Handle("/edge_groups",
		bouncer.AdminAccess(bouncer.EdgeComputeOperation(httperror.LoggerHandler(h.edgeGroupList)))).Methods(http.MethodGet)
	h.Handle("/edge_groups/{id}",
		bouncer.AdminAccess(bouncer.EdgeComputeOperation(httperror.LoggerHandler(h.edgeGroupInspect)))).Methods(http.MethodGet)
	h.Handle("/edge_groups/{id}",
		bouncer.AdminAccess(bouncer.EdgeComputeOperation(httperror.LoggerHandler(h.edgeGroupUpdate)))).Methods(http.MethodPut)
	h.Handle("/edge_groups/{id}",
		bouncer.AdminAccess(bouncer.EdgeComputeOperation(httperror.LoggerHandler(h.edgeGroupDelete)))).Methods(http.MethodDelete)

	return h
}
