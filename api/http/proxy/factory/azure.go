package factory

import (
	"net/http"
	"net/url"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	"github.com/opendocking/opendocking/api/http/proxy/factory/azure"
)

func newAzureProxy(endpoint *portainer.Endpoint, dataStore dataservices.DataStore) (http.Handler, error) {
	remoteURL, err := url.Parse(azureAPIBaseURL)
	if err != nil {
		return nil, err
	}

	proxy := NewSingleHostReverseProxyWithHostHeader(remoteURL)
	proxy.Transport = azure.NewTransport(&endpoint.AzureCredentials, dataStore, endpoint)
	return proxy, nil
}
