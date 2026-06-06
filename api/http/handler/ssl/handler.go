package ssl

import (
	"net/http"

	"github.com/opendocking/opendocking/api/http/security"
	"github.com/opendocking/opendocking/api/internal/ssl"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"

	"github.com/gorilla/mux"
)

// Handler is the HTTP handler used to handle MOTD operations.
type Handler struct {
	*mux.Router
	SSLService *ssl.Service
}

// NewHandler returns a new Handler
func NewHandler(bouncer security.BouncerService) *Handler {
	h := &Handler{
		Router: mux.NewRouter(),
	}
	h.Handle("/ssl",
		bouncer.AdminAccess(httperror.LoggerHandler(h.sslInspect))).Methods(http.MethodGet)
	h.Handle("/ssl",
		bouncer.AdminAccess(httperror.LoggerHandler(h.sslUpdate))).Methods(http.MethodPut)

	return h
}
