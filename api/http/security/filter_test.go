package security

import (
	"testing"

	portainer "github.com/opendocking/opendocking/api"
)

func TestFilterEndpointsPanic(t *testing.T) {
	t.Parallel()
	endpoints := []portainer.Endpoint{{ID: 1}}
	groups := []portainer.EndpointGroup{}
	context := &RestrictedRequestContext{}

	FilterEndpoints(endpoints, groups, context)
}
