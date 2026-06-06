package system

import (
	"net/http"

	portainer "github.com/opendocking/opendocking/api"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"
	"github.com/opendocking/opendocking/pkg/libhttp/response"

	"github.com/rs/zerolog/log"
)

type status struct {
	*portainer.Status
}

// @id systemStatus
// @summary Check Portainer status
// @description Retrieve Portainer status
// @description **Access policy**: public
// @tags system
// @produce json
// @success 200 {object} status "Success"
// @router /system/status [get]
func (handler *Handler) systemStatus(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	return response.JSON(w, &status{
		Status: handler.status,
	})
}

// swagger docs for deprecated route:
// @id StatusInspect
// @summary Check Portainer status
// @deprecated
// @description Deprecated: use the `/system/status` endpoint instead.
// @description Retrieve Portainer status
// @description **Access policy**: public
// @tags status
// @produce json
// @success 200 {object} status "Success"
// @router /status [get]
func (handler *Handler) statusInspectDeprecated(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	log.Warn().Msg("The /status endpoint is deprecated and will be removed in a future version of OpenDocking. Please use the /system/status endpoint instead.")

	return handler.systemStatus(w, r)
}
