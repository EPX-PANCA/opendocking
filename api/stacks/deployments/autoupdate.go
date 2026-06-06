package deployments

import (
	"context"
	"time"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	"github.com/opendocking/opendocking/api/scheduler"
	httperror "github.com/opendocking/opendocking/pkg/libhttp/error"

	"github.com/rs/zerolog/log"
)

func StartAutoupdate(ctx context.Context, stackID portainer.StackID, interval string, scheduler *scheduler.Scheduler, stackDeployer StackDeployer, datastore dataservices.DataStore, gitService portainer.GitService) (jobID string, e *httperror.HandlerError) {
	d, err := time.ParseDuration(interval)
	if err != nil {
		return "", httperror.BadRequest("Unable to parse stack's auto update interval", err)
	}

	jobID = scheduler.StartJobEvery(d, func() error {
		return RedeployWhenChanged(ctx, stackID, stackDeployer, datastore, gitService)
	})

	return jobID, nil
}

func StopAutoupdate(stackID portainer.StackID, jobID string, scheduler *scheduler.Scheduler) {
	if jobID == "" {
		return
	}

	if err := scheduler.StopJob(jobID); err != nil {
		log.Warn().Int("stack_id", int(stackID)).Msg("could not stop the job for the stack")
	}
}
