package stackbuilders

import (
	"context"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	"github.com/opendocking/opendocking/api/filesystem"
	"github.com/opendocking/opendocking/api/http/security"
	"github.com/opendocking/opendocking/api/stacks/deployments"
)

type SwarmStackFileBuilder struct {
	StackBuilder
	SecurityContext *security.RestrictedRequestContext
}

// CreateSwarmStackFileBuilder creates a builder for swarm stacks deployed from a file (either uploaded or provided as text content).
func CreateSwarmStackFileBuilder(securityContext *security.RestrictedRequestContext,
	dataStore dataservices.DataStore,
	fileService portainer.FileService,
	stackDeployer deployments.StackDeployer) *SwarmStackFileBuilder {

	return &SwarmStackFileBuilder{
		StackBuilder:    CreateStackBuilder(dataStore, fileService, stackDeployer),
		SecurityContext: securityContext,
	}
}

func (b *SwarmStackFileBuilder) prepare(_ context.Context, payload *StackPayload, userID portainer.UserID) error {
	b.stack.Name = payload.Name
	b.stack.Type = portainer.DockerSwarmStack
	b.stack.SwarmID = payload.SwarmID
	b.stack.EntryPoint = filesystem.ComposeFileDefaultName
	b.stack.Env = payload.Env
	b.stack.FromAppTemplate = payload.FromAppTemplate

	if err := b.initCreatedBy(userID); err != nil {
		return err
	}

	return b.storeStackFile(payload.StackFileContent)
}

func (b *SwarmStackFileBuilder) deploy(ctx context.Context, endpoint *portainer.Endpoint) error {
	if err := b.initSwarmDeployment(b.SecurityContext, endpoint); err != nil {
		return err
	}

	return b.deploymentConfiger.Deploy(ctx)
}
