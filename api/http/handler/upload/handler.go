package upload

import (
	"net/http"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/http/security"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"

	"github.com/gorilla/mux"
)

// Handler is the HTTP handler used to handle upload operations.
type Handler struct {
	*mux.Router
	FileService portainer.FileService
}

// NewHandler creates a handler to manage upload operations.
func NewHandler(bouncer security.BouncerService) *Handler {
	h := &Handler{
		Router: mux.NewRouter(),
	}
	h.Handle("/upload/tls/{certificate:(?:ca|cert|key)}",
		bouncer.AdminAccess(httperror.LoggerHandler(h.uploadTLS))).Methods(http.MethodPost)
	return h
}
