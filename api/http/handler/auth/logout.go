package auth

import (
	"net/http"
	"strconv"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	"github.com/opendocking/opendocking/api/http/security"
	"github.com/opendocking/opendocking/api/logoutcontext"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"
	"github.com/opendocking/opendocking/pkg/libhttp/response"
)

// @id Logout
// @summary Logout
// @description **Access policy**: public
// @security ApiKeyAuth
// @security jwt
// @tags auth
// @success 204 "Success"
// @failure 500 "Server error"
// @router /auth/logout [post]
func (handler *Handler) logout(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	tokenData, _ := handler.bouncer.CookieAuthLookup(r)

	if tokenData != nil {
		handler.KubernetesTokenCacheManager.RemoveUserFromCache(tokenData.ID)
		handler.KubernetesClientFactory.ClearUserClientCache(strconv.Itoa(int(tokenData.ID)))
		logoutcontext.Cancel(tokenData.Token)
		handler.bouncer.RevokeJWT(tokenData.Token)
	}

	var settings *portainer.Settings
	if err := handler.DataStore.ViewTx(func(tx dataservices.DataStoreTx) error {
		var err error
		settings, err = tx.Settings().Settings()
		return err
	}); err != nil {
		return httperror.InternalServerError("Unable to retrieve settings from the database", err)
	}

	security.RemoveAuthCookie(w, handler.isSecureCookie(r, settings.ForceSecureCookies))

	return response.Empty(w)
}
