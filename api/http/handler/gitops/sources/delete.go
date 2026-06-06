package sources

import (
	"errors"
	"net/http"
	"slices"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	dserrors "github.com/opendocking/opendocking/api/dataservices/errors"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"
	"github.com/opendocking/opendocking/pkg/libhttp/request"
	"github.com/opendocking/opendocking/pkg/libhttp/response"
)

var ErrSourceInUse = errors.New("source is used by one or more workflows or custom templates")

// @id GitOpsSourcesDelete
// @summary Delete a source
// @description Deletes an existing GitOps source. Returns 409 if the source is referenced by any workflow or custom template.
// @description **Access policy**: admin
// @tags gitops
// @security ApiKeyAuth
// @security jwt
// @param id path int true "Source identifier"
// @success 204 "Source deleted"
// @failure 400 "Invalid request"
// @failure 403 "Access denied"
// @failure 404 "Source not found"
// @failure 409 "Source is in use by one or more workflows or custom templates"
// @failure 500 "Server error"
// @router /gitops/sources/{id} [delete]
func (h *Handler) sourceDelete(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	sourceID, err := request.RetrieveNumericRouteVariableValue(r, "id")
	if err != nil {
		return httperror.BadRequest("Invalid source identifier route variable", err)
	}

	if err := h.dataStore.UpdateTx(func(tx dataservices.DataStoreTx) error {
		if exists, err := tx.Source().Exists(portainer.SourceID(sourceID)); err != nil {
			return err
		} else if !exists {
			return dserrors.ErrObjectNotFound
		}

		workflows, err := tx.Workflow().ReadAll()
		if err != nil {
			return err
		}

		for _, wf := range workflows {
			if slices.ContainsFunc(wf.Artifacts, func(as portainer.ArtifactSources) bool {
				return slices.Contains(as.SourceIDs, portainer.SourceID(sourceID))
			}) {
				return ErrSourceInUse
			}
		}

		templates, err := tx.CustomTemplate().ReadAll(func(t portainer.CustomTemplate) bool {
			return t.ArtifactSources != nil && slices.Contains(t.ArtifactSources.SourceIDs, portainer.SourceID(sourceID))
		})
		if err != nil {
			return err
		}

		if len(templates) > 0 {
			return ErrSourceInUse
		}

		return tx.Source().Delete(portainer.SourceID(sourceID))
	}); h.dataStore.IsErrObjectNotFound(err) {
		return httperror.NotFound("Unable to find a source with the specified identifier", err)
	} else if errors.Is(err, ErrSourceInUse) {
		return httperror.Conflict("Source is used by one or more workflows or custom templates", err)
	} else if err != nil {
		return httperror.InternalServerError("Unable to delete source", err)
	}

	return response.Empty(w)
}
