package uac

import (
	"testing"

	"github.com/docker/docker/api/types/swarm"
	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	"github.com/opendocking/opendocking/api/datastore"
	"github.com/opendocking/opendocking/api/docker/consts"
	"github.com/opendocking/opendocking/api/internal/authorization"
	"github.com/opendocking/opendocking/api/stacks/stackutils"
	"github.com/stretchr/testify/require"
)

func TestSecretResourceControlGetter(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	ok, store := datastore.MustNewTestStore(t, true, false)
	is.True(ok)
	is.NotNil(store)

	envID := portainer.EndpointID(1)
	secretID := "secret"
	stackName := "stack"
	stackRCID := stackutils.ResourceControlID(envID, stackName)
	serviceID := "service"

	is.NoError(store.UpdateTx(func(tx dataservices.DataStoreTx) error {
		is.NoError(tx.ResourceControl().Create(authorization.NewPublicResourceControl(secretID, portainer.SecretResourceControl)))
		is.NoError(tx.ResourceControl().Create(authorization.NewPublicResourceControl(stackRCID, portainer.StackResourceControl)))
		is.NoError(tx.ResourceControl().Create(authorization.NewPublicResourceControl(serviceID, portainer.ServiceResourceControl)))
		return nil
	}))

	is.NoError(store.ViewTx(func(tx dataservices.DataStoreTx) error {
		// by direct ID
		rc, err := SecretResourceControlGetter(tx, envID)(swarm.Secret{ID: secretID})
		is.NoError(err)
		is.NotNil(rc)
		is.Equal(secretID, rc.ResourceID)

		// by compose stack label
		rc, err = SecretResourceControlGetter(tx, envID)(
			swarm.Secret{ID: "unknown", Spec: swarm.SecretSpec{Annotations: swarm.Annotations{Labels: map[string]string{consts.ComposeStackNameLabel: stackName}}}},
		)
		is.NoError(err)
		is.NotNil(rc)
		is.Equal(stackRCID, rc.ResourceID)

		// by swarm stack label
		rc, err = SecretResourceControlGetter(tx, envID)(
			swarm.Secret{ID: "unknown", Spec: swarm.SecretSpec{Annotations: swarm.Annotations{Labels: map[string]string{consts.SwarmStackNameLabel: stackName}}}},
		)
		is.NoError(err)
		is.NotNil(rc)
		is.Equal(stackRCID, rc.ResourceID)

		// by service ID
		rc, err = SecretResourceControlGetter(tx, envID)(
			swarm.Secret{ID: "unknown", Spec: swarm.SecretSpec{Annotations: swarm.Annotations{Labels: map[string]string{consts.SwarmServiceIDLabel: serviceID}}}},
		)
		is.NoError(err)
		is.NotNil(rc)
		is.Equal(serviceID, rc.ResourceID)

		return nil
	}))

}
