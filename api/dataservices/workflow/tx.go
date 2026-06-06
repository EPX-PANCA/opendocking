package workflow

import (
	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
)

type ServiceTx struct {
	dataservices.BaseDataServiceTx[portainer.Workflow, portainer.WorkflowID]
}

func (service ServiceTx) Create(workflow *portainer.Workflow) error {
	return service.Tx.CreateObject(
		BucketName,
		func(id uint64) (int, any) {
			workflow.ID = portainer.WorkflowID(id)
			return int(workflow.ID), workflow
		},
	)
}
