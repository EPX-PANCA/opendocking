package auth

import (
	"net/http"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	"github.com/opendocking/opendocking/api/http/proxy"
	"github.com/opendocking/opendocking/api/http/proxy/factory/kubernetes"
	"github.com/opendocking/opendocking/api/http/security"
	"github.com/opendocking/opendocking/api/kubernetes/cli"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"

	"github.com/gorilla/mux"
)

// Handler is the HTTP handler used to handle authentication operations.
type Handler struct {
	*mux.Router
	DataStore                   dataservices.DataStore
	CryptoService               portainer.CryptoService
	JWTService                  portainer.JWTService
	LDAPService                 portainer.LDAPService
	OAuthService                portainer.OAuthService
	ProxyManager                *proxy.Manager
	KubernetesTokenCacheManager *kubernetes.TokenCacheManager
	KubernetesClientFactory     *cli.ClientFactory
	passwordStrengthChecker     security.PasswordStrengthChecker
	bouncer                     security.BouncerService
}

// NewHandler creates a handler to manage authentication operations.
func NewHandler(bouncer security.BouncerService, rateLimiter *security.RateLimiter, passwordStrengthChecker security.PasswordStrengthChecker, kubernetesClientFactory *cli.ClientFactory) *Handler {
	h := &Handler{
		Router:                  mux.NewRouter(),
		passwordStrengthChecker: passwordStrengthChecker,
		bouncer:                 bouncer,
		KubernetesClientFactory: kubernetesClientFactory,
	}

	h.Handle("/auth/oauth/validate",
		rateLimiter.LimitAccess(bouncer.PublicAccess(httperror.LoggerHandler(h.validateOAuth)))).Methods(http.MethodPost)
	h.Handle("/auth",
		rateLimiter.LimitAccess(bouncer.PublicAccess(httperror.LoggerHandler(h.authenticate)))).Methods(http.MethodPost)
	h.Handle("/auth/logout",
		bouncer.PublicAccess(httperror.LoggerHandler(h.logout))).Methods(http.MethodPost)

	return h
}
