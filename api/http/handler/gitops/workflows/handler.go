package workflows

import (
	"net/http"
	"time"

	gocache "github.com/patrickmn/go-cache"
	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	"github.com/opendocking/opendocking/api/kubernetes/cli"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"

	"github.com/gorilla/mux"
)

const (
	cacheTTL             = 30 * time.Second
	cacheCleanupInterval = 10 * time.Minute
)

type Handler struct {
	*mux.Router
	dataStore  dataservices.DataStore
	gitService portainer.GitService
	cache      *gocache.Cache
	k8sFactory *cli.ClientFactory
}

func NewHandler(dataStore dataservices.DataStore, gitService portainer.GitService, k8sFactory *cli.ClientFactory) *Handler {
	h := &Handler{
		Router:     mux.NewRouter(),
		dataStore:  dataStore,
		gitService: gitService,
		cache:      gocache.New(cacheTTL, cacheCleanupInterval),
		k8sFactory: k8sFactory,
	}

	h.Handle("/gitops/workflows", httperror.LoggerHandler(h.list)).Methods(http.MethodGet)
	h.Handle("/gitops/workflows/summary", httperror.LoggerHandler(h.summary)).Methods(http.MethodGet)

	return h
}
