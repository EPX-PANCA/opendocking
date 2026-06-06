package gitops

import (
	"net/http"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	"github.com/opendocking/opendocking/api/http/security"
	"github.com/opendocking/opendocking/api/kubernetes/cli"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"

	"github.com/gorilla/mux"

	"github.com/opendocking/opendocking/api/http/handler/gitops/sources"
	"github.com/opendocking/opendocking/api/http/handler/gitops/workflows"
)

// Handler is the HTTP handler used to handle git repo operation
type Handler struct {
	*mux.Router
	dataStore   dataservices.DataStore
	gitService  portainer.GitService
	fileService portainer.FileService
}

func NewHandler(bouncer security.BouncerService, dataStore dataservices.DataStore, gitService portainer.GitService, fileService portainer.FileService, k8sFactory *cli.ClientFactory) *Handler {
	h := &Handler{
		Router:      mux.NewRouter(),
		dataStore:   dataStore,
		gitService:  gitService,
		fileService: fileService,
	}

	authenticatedRouter := h.NewRoute().Subrouter()
	authenticatedRouter.Use(bouncer.AuthenticatedAccess)

	authenticatedRouter.Handle("/gitops/repo/file/preview", httperror.LoggerHandler(h.gitOperationRepoFilePreview)).Methods(http.MethodPost)

	workflowsHandler := workflows.NewHandler(dataStore, gitService, k8sFactory)
	authenticatedRouter.PathPrefix("/gitops/workflows").Handler(workflowsHandler)

	sourcesHandler := sources.NewHandler(bouncer, dataStore, gitService, k8sFactory)
	authenticatedRouter.PathPrefix("/gitops/sources").Handler(sourcesHandler)

	return h
}
