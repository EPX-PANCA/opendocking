package websocket

import portainer "github.com/opendocking/opendocking/api"

type webSocketRequestParams struct {
	ID       string
	nodeName string
	endpoint *portainer.Endpoint
	token    string
}
