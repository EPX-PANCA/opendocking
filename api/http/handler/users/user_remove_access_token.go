package users

import (
	"errors"
	"net/http"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/apikey"
	httperrors "github.com/opendocking/opendocking/api/http/errors"
	"github.com/opendocking/opendocking/api/http/security"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"
	"github.com/opendocking/opendocking/pkg/libhttp/request"
	"github.com/opendocking/opendocking/pkg/libhttp/response"
)

// @id UserRemoveAPIKey
// @summary Remove an api-key associated to a user
// @description Remove an api-key associated to a user..
// @description Only the calling user or admin can remove api-key.
// @description **Access policy**: authenticated
// @tags users
// @security ApiKeyAuth
// @security jwt
// @param id path int true "User identifier"
// @param keyID path int true "Api Key identifier"
// @success 204 "Success"
// @failure 400 "Invalid request"
// @failure 403 "Permission denied"
// @failure 404 "Not found"
// @failure 500 "Server error"
// @router /users/{id}/tokens/{keyID} [delete]
func (handler *Handler) userRemoveAccessToken(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	userID, err := request.RetrieveNumericRouteVariableValue(r, "id")
	if err != nil {
		return httperror.BadRequest("Invalid user identifier route variable", err)
	}

	apiKeyID, err := request.RetrieveNumericRouteVariableValue(r, "keyID")
	if err != nil {
		return httperror.BadRequest("Invalid api-key identifier route variable", err)
	}

	tokenData, err := security.RetrieveTokenData(r)
	if err != nil {
		return httperror.InternalServerError("Unable to retrieve user authentication token", err)
	}
	if tokenData.Role != portainer.AdministratorRole && tokenData.ID != portainer.UserID(userID) {
		return httperror.Forbidden("Permission denied to get user access tokens", httperrors.ErrUnauthorized)
	}

	_, err = handler.DataStore.User().Read(portainer.UserID(userID))
	if err != nil {
		if handler.DataStore.IsErrObjectNotFound(err) {
			return httperror.NotFound("Unable to find a user with the specified identifier inside the database", err)
		}
		return httperror.InternalServerError("Unable to find a user with the specified identifier inside the database", err)
	}

	// check if the key exists and the key belongs to the user
	apiKey, err := handler.apiKeyService.GetAPIKey(portainer.APIKeyID(apiKeyID))
	if err != nil {
		return httperror.InternalServerError("API Key not found", err)
	}
	if apiKey.UserID != portainer.UserID(userID) {
		return httperror.Forbidden("Permission denied to remove api-key", httperrors.ErrUnauthorized)
	}

	err = handler.apiKeyService.DeleteAPIKey(portainer.APIKeyID(apiKeyID))
	if err != nil {
		if errors.Is(err, apikey.ErrInvalidAPIKey) {
			return httperror.NotFound("Unable to find an api-key with the specified identifier inside the database", err)
		}
		return httperror.InternalServerError("Unable to remove the api-key from the user", err)
	}

	return response.Empty(w)
}
