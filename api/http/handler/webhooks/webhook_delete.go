package webhooks

import (
	"errors"
	"net/http"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/http/security"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"
	"github.com/opendocking/opendocking/pkg/libhttp/request"
	"github.com/opendocking/opendocking/pkg/libhttp/response"
)

// @summary Delete a webhook
// @description **Access policy**: authenticated
// @security ApiKeyAuth
// @security jwt
// @tags webhooks
// @param id path int true "Webhook id"
// @success 202 "Webhook deleted"
// @failure 400
// @failure 500
// @router /webhooks/{id} [delete]
func (handler *Handler) webhookDelete(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	id, err := request.RetrieveNumericRouteVariableValue(r, "id")
	if err != nil {
		return httperror.BadRequest("Invalid webhook id", err)
	}

	securityContext, err := security.RetrieveRestrictedRequestContext(r)
	if err != nil {
		return httperror.InternalServerError("Unable to retrieve user info from request context", err)
	}

	if !securityContext.IsAdmin {
		return httperror.Forbidden("Not authorized to delete a webhook", errors.New("not authorized to delete a webhook"))
	}

	err = handler.DataStore.Webhook().Delete(portainer.WebhookID(id))
	if err != nil {
		return httperror.InternalServerError("Unable to remove the webhook from the database", err)
	}

	return response.Empty(w)
}
