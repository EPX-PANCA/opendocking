package docker

import (
	portainer "github.com/opendocking/opendocking/api"
	dockerclient "github.com/opendocking/opendocking/api/docker/client"
	"github.com/opendocking/opendocking/api/logs"
	"github.com/opendocking/opendocking/pkg/snapshot"
)

// Snapshotter represents a service used to create environment(endpoint) snapshots
type Snapshotter struct {
	clientFactory *dockerclient.ClientFactory
}

// NewSnapshotter returns a new Snapshotter instance
func NewSnapshotter(clientFactory *dockerclient.ClientFactory) *Snapshotter {
	return &Snapshotter{
		clientFactory: clientFactory,
	}
}

// CreateSnapshot creates a snapshot of a specific Docker environment(endpoint)
func (snapshotter *Snapshotter) CreateSnapshot(endpoint *portainer.Endpoint) (*portainer.DockerSnapshot, error) {
	cli, err := snapshotter.clientFactory.CreateClient(endpoint, "", nil)
	if err != nil {
		return nil, err
	}
	defer logs.CloseAndLogErr(cli)

	return snapshot.CreateDockerSnapshot(cli)
}
