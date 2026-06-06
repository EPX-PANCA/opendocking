package customtemplates

import (
	"net/http"
	"strconv"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	httperrors "github.com/opendocking/opendocking/api/http/errors"
	"github.com/opendocking/opendocking/api/http/security"
	"github.com/opendocking/opendocking/api/internal/authorization"
	"github.com/opendocking/opendocking/api/slicesx"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"
	"github.com/opendocking/opendocking/pkg/libhttp/request"
	"github.com/opendocking/opendocking/pkg/libhttp/response"
)

// @id CustomTemplateInspect
// @summary Inspect a custom template
// @description Retrieve details about a template.
// @description **Access policy**: authenticated
// @tags custom_templates
// @security ApiKeyAuth
// @security jwt
// @produce json
// @param id path int true "Template identifier"
// @success 200 {object} portainer.CustomTemplate "Success"
// @failure 400 "Invalid request"
// @failure 404 "Template not found"
// @failure 500 "Server error"
// @router /custom_templates/{id} [get]
func (handler *Handler) customTemplateInspect(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	customTemplateID, err := request.RetrieveNumericRouteVariableValue(r, "id")
	if err != nil {
		return httperror.BadRequest("Invalid Custom template identifier route variable", err)
	}

	var customTemplate *portainer.CustomTemplate
	err = handler.DataStore.ViewTx(func(tx dataservices.DataStoreTx) error {
		customTemplate, err = tx.CustomTemplate().Read(portainer.CustomTemplateID(customTemplateID))
		if tx.IsErrObjectNotFound(err) {
			return httperror.NotFound("Unable to find a custom template with the specified identifier inside the database", err)
		} else if err != nil {
			return httperror.InternalServerError("Unable to find a custom template with the specified identifier inside the database", err)
		}

		resourceControl, err := tx.ResourceControl().ResourceControlByResourceIDAndType(strconv.Itoa(customTemplateID), portainer.CustomTemplateResourceControl)
		if err != nil {
			return httperror.InternalServerError("Unable to retrieve a resource control associated to the custom template", err)
		}

		securityContext, err := security.RetrieveRestrictedRequestContext(r)
		if err != nil {
			return httperror.InternalServerError("Unable to retrieve user info from request context", err)
		}

		canEdit := userCanEditTemplate(customTemplate, securityContext)
		hasAccess := false

		if resourceControl != nil {
			customTemplate.ResourceControl = resourceControl

			teamIDs := slicesx.Map(securityContext.UserMemberships, func(m portainer.TeamMembership) portainer.TeamID {
				return m.TeamID
			})

			hasAccess = authorization.UserCanAccessResource(securityContext.UserID, teamIDs, resourceControl)

		}

		if !canEdit && !hasAccess {
			return httperror.Forbidden("Access denied to resource", httperrors.ErrResourceAccessDenied)
		}

		populateGitConfig(tx, customTemplate)

		return nil
	})

	return response.TxResponse(w, customTemplate, err)
}
