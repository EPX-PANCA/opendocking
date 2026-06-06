package ldap

import (
	"net/http"

	portainer "github.com/opendocking/opendocking/api"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"
	"github.com/opendocking/opendocking/pkg/libhttp/request"
	"github.com/opendocking/opendocking/pkg/libhttp/response"
)

type checkPayload struct {
	LDAPSettings portainer.LDAPSettings
}

func (payload *checkPayload) Validate(r *http.Request) error {
	return nil
}

// @id LDAPCheck
// @summary Test LDAP connectivity
// @description Test LDAP connectivity using LDAP details
// @description **Access policy**: administrator
// @tags ldap
// @security ApiKeyAuth
// @security jwt
// @accept json
// @param body body checkPayload true "details"
// @success 204 "Success"
// @failure 400 "Invalid request"
// @failure 500 "Server error"
// @router /ldap/check [post]
func (handler *Handler) ldapCheck(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	var payload checkPayload
	err := request.DecodeAndValidateJSONPayload(r, &payload)
	if err != nil {
		return httperror.BadRequest("Invalid request payload", err)
	}

	settings := &payload.LDAPSettings

	err = handler.prefillSettings(settings)
	if err != nil {
		return httperror.InternalServerError("Unable to fetch default settings", err)
	}

	err = handler.LDAPService.TestConnectivity(settings)
	if err != nil {
		return httperror.InternalServerError("Unable to connect to LDAP server", err)
	}

	return response.Empty(w)
}
