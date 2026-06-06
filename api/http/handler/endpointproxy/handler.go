package endpointproxy

import (
	"github.com/gorilla/mux"
	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	"github.com/opendocking/opendocking/api/http/proxy"
	"github.com/opendocking/opendocking/api/http/security"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"
)

// Handler is the HTTP handler used to proxy requests to external APIs.
type Handler struct {
	*mux.Router
	DataStore            dataservices.DataStore
	requestBouncer       security.BouncerService
	ProxyManager         *proxy.Manager
	ReverseTunnelService portainer.ReverseTunnelService
}

// NewHandler creates a handler to proxy requests to external APIs.
func NewHandler(bouncer security.BouncerService) *Handler {
	h := &Handler{
		Router:         mux.NewRouter(),
		requestBouncer: bouncer,
	}
	h.PathPrefix("/{id}/azure").Handler(
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.proxyRequestsToAzureAPI)))
	h.PathPrefix("/{id}/docker").Handler(
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.proxyRequestsToDockerAPI)))
	h.PathPrefix("/{id}/kubernetes").Handler(
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.proxyRequestsToKubernetesAPI)))
	h.PathPrefix("/{id}/agent/docker").Handler(
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.proxyRequestsToDockerAPI)))
	h.PathPrefix("/{id}/agent/kubernetes").Handler(
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.proxyRequestsToKubernetesAPI)))
	h.PathPrefix("/{id}/agent/host").Handler(httperror.LoggerHandler(h.proxyRequestsToAgentHostAPI))
	return h
}
