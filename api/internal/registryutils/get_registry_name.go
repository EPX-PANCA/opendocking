package registryutils

import (
	"strconv"

	portainer "github.com/opendocking/opendocking/api"
)

func RegistrySecretName(registryID portainer.RegistryID) string {
	return "registry-" + strconv.Itoa(int(registryID))
}
