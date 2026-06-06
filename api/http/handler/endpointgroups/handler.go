package endpointgroups

import (
	"net/http"

	"github.com/opendocking/opendocking/api/dataservices"
	"github.com/opendocking/opendocking/api/http/security"
	"github.com/opendocking/opendocking/api/internal/authorization"
	"github.com/opendocking/opendocking/api/pendingactions"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"

	"github.com/gorilla/mux"
)

// Handler is the HTTP handler used to handle environment(endpoint) group operations.
type Handler struct {
	*mux.Router
	AuthorizationService  *authorization.Service
	DataStore             dataservices.DataStore
	PendingActionsService *pendingactions.PendingActionsService
}

// NewHandler creates a handler to manage environment(endpoint) group operations.
func NewHandler(bouncer security.BouncerService) *Handler {
	h := &Handler{
		Router: mux.NewRouter(),
	}
	h.Handle("/endpoint_groups",
		bouncer.AdminAccess(httperror.LoggerHandler(h.endpointGroupCreate))).Methods(http.MethodPost)
	h.Handle("/endpoint_groups",
		bouncer.RestrictedAccess(httperror.LoggerHandler(h.endpointGroupList))).Methods(http.MethodGet)
	h.Handle("/endpoint_groups/{id}",
		bouncer.AdminAccess(httperror.LoggerHandler(h.endpointGroupInspect))).Methods(http.MethodGet)
	h.Handle("/endpoint_groups/{id}",
		bouncer.AdminAccess(httperror.LoggerHandler(h.endpointGroupUpdate))).Methods(http.MethodPut)
	h.Handle("/endpoint_groups/{id}",
		bouncer.AdminAccess(httperror.LoggerHandler(h.endpointGroupDelete))).Methods(http.MethodDelete)
	h.Handle("/endpoint_groups/{id}/endpoints/{endpointId}",
		bouncer.AdminAccess(httperror.LoggerHandler(h.endpointGroupAddEndpoint))).Methods(http.MethodPut)
	h.Handle("/endpoint_groups/{id}/endpoints/{endpointId}",
		bouncer.AdminAccess(httperror.LoggerHandler(h.endpointGroupDeleteEndpoint))).Methods(http.MethodDelete)
	return h
}
