package source

import (
	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
)

type ServiceTx struct {
	dataservices.BaseDataServiceTx[portainer.Source, portainer.SourceID]
}

// Create creates a new source.
func (service ServiceTx) Create(source *portainer.Source) error {
	return service.Tx.CreateObject(
		BucketName,
		func(id uint64) (int, any) {
			source.ID = portainer.SourceID(id)
			return int(source.ID), source
		},
	)
}
