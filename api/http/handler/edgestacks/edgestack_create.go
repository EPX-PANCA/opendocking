package edgestacks

import (
	"net/http"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	httperrors "github.com/opendocking/opendocking/api/http/errors"
	"github.com/opendocking/opendocking/api/http/security"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"
	"github.com/opendocking/opendocking/pkg/libhttp/request"
	"github.com/opendocking/opendocking/pkg/libhttp/response"
)

func (handler *Handler) edgeStackCreate(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	method, err := request.RetrieveRouteVariableValue(r, "method")
	if err != nil {
		return httperror.BadRequest("Invalid query parameter: method", err)
	}

	dryrun, _ := request.RetrieveBooleanQueryParameter(r, "dryrun", true)

	tokenData, err := security.RetrieveTokenData(r)
	if err != nil {
		return httperror.InternalServerError("Unable to retrieve user details from authentication token", err)
	}

	var edgeStack *portainer.EdgeStack
	if err := handler.DataStore.UpdateTx(func(tx dataservices.DataStoreTx) error {
		edgeStack, err = handler.createSwarmStack(tx, method, dryrun, tokenData, r)
		return err
	}); err != nil {
		switch {
		case httperrors.IsInvalidPayloadError(err):
			return httperror.BadRequest("Invalid payload", err)
		case httperrors.IsConflictError(err):
			return httperror.Conflict(err.Error(), err)
		default:
			return httperror.InternalServerError("Unable to create Edge stack", err)
		}
	}

	return response.JSON(w, edgeStack)
}

func (handler *Handler) createSwarmStack(tx dataservices.DataStoreTx, method string, dryrun bool, tokenData *portainer.TokenData, r *http.Request) (*portainer.EdgeStack, error) {
	switch method {
	case "string":
		return handler.createEdgeStackFromFileContent(r, tx, tokenData, dryrun)
	case "repository":
		return handler.createEdgeStackFromGitRepository(r, tx, tokenData, dryrun)
	case "file":
		return handler.createEdgeStackFromFileUpload(r, tx, tokenData, dryrun)
	}

	return nil, httperrors.NewInvalidPayloadError("Invalid value for query parameter: method. Value must be one of: string, repository or file")
}
