package system

import (
	"net/http"

	portainer "github.com/opendocking/opendocking/api"
	statusutil "github.com/opendocking/opendocking/api/internal/nodes"
	"github.com/opendocking/opendocking/api/internal/snapshot"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"
	"github.com/opendocking/opendocking/pkg/libhttp/response"
)

type nodesCountResponse struct {
	Nodes int `json:"nodes"`
}

// @id systemNodesCount
// @summary Retrieve the count of nodes
// @description **Access policy**: authenticated
// @security ApiKeyAuth
// @security jwt
// @tags system
// @produce json
// @success 200 {object} nodesCountResponse "Success"
// @failure 500 "Server error"
// @router /system/nodes [get]
func (handler *Handler) systemNodesCount(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	endpoints, err := handler.dataStore.Endpoint().Endpoints()
	if err != nil {
		return httperror.InternalServerError("Failed to get environment list", err)
	}

	var nodes int

	for _, endpoint := range endpoints {
		if err := snapshot.FillSnapshotData(handler.dataStore, &endpoint, false); err != nil {
			return httperror.InternalServerError("Unable to add snapshot data", err)
		}

		nodes += statusutil.NodesCount([]portainer.Endpoint{endpoint})
	}

	return response.JSON(w, &nodesCountResponse{Nodes: nodes})
}
