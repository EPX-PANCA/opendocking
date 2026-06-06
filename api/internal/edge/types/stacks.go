package types

import portainer "github.com/opendocking/opendocking/api"

type StoreManifestFunc func(stackFolder string, relatedEndpointIds []portainer.EndpointID) (composePath, manifestPath, projectPath string, err error)
