package edgegroups

import (
	"net/http"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	"github.com/opendocking/opendocking/api/roar"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"
	"github.com/opendocking/opendocking/pkg/libhttp/request"
	"github.com/opendocking/opendocking/pkg/libhttp/response"
)

// @id EdgeGroupInspect
// @summary Inspects an EdgeGroup
// @description **Access policy**: administrator
// @tags edge_groups
// @security ApiKeyAuth
// @security jwt
// @produce json
// @param id path int true "EdgeGroup Id"
// @success 200 {object} portainer.EdgeGroup
// @failure 503 "Edge compute features are disabled"
// @failure 500
// @router /edge_groups/{id} [get]
func (handler *Handler) edgeGroupInspect(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	edgeGroupID, err := request.RetrieveNumericRouteVariableValue(r, "id")
	if err != nil {
		return httperror.BadRequest("Invalid Edge group identifier route variable", err)
	}

	var shadowEdgeGroup shadowedEdgeGroup
	err = handler.DataStore.ViewTx(func(tx dataservices.DataStoreTx) error {
		edgeGroup, err := getEdgeGroup(tx, portainer.EdgeGroupID(edgeGroupID))
		if err != nil {
			return err
		}

		edgeGroup.Endpoints = edgeGroup.EndpointIDs.ToSlice()

		shadowEdgeGroup = shadowedEdgeGroup{EdgeGroup: *edgeGroup}

		return nil
	})

	return response.TxResponse(w, shadowEdgeGroup, err)
}

func getEdgeGroup(tx dataservices.DataStoreTx, ID portainer.EdgeGroupID) (*portainer.EdgeGroup, error) {
	edgeGroup, err := tx.EdgeGroup().Read(ID)
	if tx.IsErrObjectNotFound(err) {
		return nil, httperror.NotFound("Unable to find an Edge group with the specified identifier inside the database", err)
	} else if err != nil {
		return nil, httperror.InternalServerError("Unable to find an Edge group with the specified identifier inside the database", err)
	}

	if edgeGroup.Dynamic {
		endpoints, err := GetEndpointsByTags(tx, edgeGroup.TagIDs, edgeGroup.PartialMatch)
		if err != nil {
			return nil, httperror.InternalServerError("Unable to retrieve environments and environment groups for Edge group", err)
		}

		edgeGroup.EndpointIDs = roar.FromSlice(endpoints)
	}

	return edgeGroup, err
}
