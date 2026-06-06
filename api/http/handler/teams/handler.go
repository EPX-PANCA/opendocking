package teams

import (
	"net/http"

	"github.com/opendocking/opendocking/api/dataservices"
	"github.com/opendocking/opendocking/api/http/security"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"

	"github.com/gorilla/mux"
)

// Handler is the HTTP handler used to handle team operations.
type Handler struct {
	*mux.Router
	DataStore dataservices.DataStore
}

// NewHandler creates a handler to manage team operations.
func NewHandler(bouncer security.BouncerService) *Handler {
	h := &Handler{
		Router: mux.NewRouter(),
	}

	adminRouter := h.NewRoute().Subrouter()
	adminRouter.Use(bouncer.AdminAccess)

	restrictedRouter := h.NewRoute().Subrouter()
	restrictedRouter.Use(bouncer.RestrictedAccess)

	teamLeaderRouter := h.NewRoute().Subrouter()
	teamLeaderRouter.Use(bouncer.TeamLeaderAccess)

	adminRouter.Handle("/teams", httperror.LoggerHandler(h.teamCreate)).Methods(http.MethodPost)
	restrictedRouter.Handle("/teams", httperror.LoggerHandler(h.teamList)).Methods(http.MethodGet)
	teamLeaderRouter.Handle("/teams/{id}", httperror.LoggerHandler(h.teamInspect)).Methods(http.MethodGet)
	adminRouter.Handle("/teams/{id}", httperror.LoggerHandler(h.teamUpdate)).Methods(http.MethodPut)
	adminRouter.Handle("/teams/{id}", httperror.LoggerHandler(h.teamDelete)).Methods(http.MethodDelete)
	teamLeaderRouter.Handle("/teams/{id}/memberships", httperror.LoggerHandler(h.teamMemberships)).Methods(http.MethodGet)

	return h
}
