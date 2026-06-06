package stackbuilders

import (
	"context"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	"github.com/opendocking/opendocking/api/http/security"
	"github.com/opendocking/opendocking/api/scheduler"
	"github.com/opendocking/opendocking/api/stacks/deployments"
)

type ComposeStackGitBuilder struct {
	GitMethodStackBuilder
	SecurityContext *security.RestrictedRequestContext
}

// CreateComposeStackGitBuilder creates a builder for the compose stack (docker standalone) that will be deployed by git repository method
func CreateComposeStackGitBuilder(securityContext *security.RestrictedRequestContext,
	dataStore dataservices.DataStore,
	fileService portainer.FileService,
	gitService portainer.GitService,
	scheduler *scheduler.Scheduler,
	stackDeployer deployments.StackDeployer) *ComposeStackGitBuilder {

	return &ComposeStackGitBuilder{
		GitMethodStackBuilder: GitMethodStackBuilder{
			StackBuilder: CreateStackBuilder(dataStore, fileService, stackDeployer),
			gitService:   gitService,
			scheduler:    scheduler,
		},
		SecurityContext: securityContext,
	}
}

func (b *ComposeStackGitBuilder) prepare(ctx context.Context, payload *StackPayload, userID portainer.UserID) error {
	b.stack.Name = payload.Name
	b.stack.Type = portainer.DockerComposeStack
	b.stack.EntryPoint = payload.ComposeFile
	b.stack.FromAppTemplate = payload.FromAppTemplate
	b.stack.Env = payload.Env

	return b.GitMethodStackBuilder.prepare(ctx, payload, userID)
}

func (b *ComposeStackGitBuilder) deploy(ctx context.Context, endpoint *portainer.Endpoint) error {
	if err := b.initComposeDeployment(b.SecurityContext, endpoint); err != nil {
		return err
	}

	return b.deploymentConfiger.Deploy(ctx)
}
