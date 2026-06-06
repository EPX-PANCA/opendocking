package migrator

import (
	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/internal/authorization"

	"github.com/rs/zerolog/log"
)

func (m *Migrator) migrateDBVersionToDB36() error {
	log.Info().Msg("updating user authorizations")

	return m.migrateUsersToDB36()
}

func (m *Migrator) migrateUsersToDB36() error {
	log.Info().Msg("updating user authorizations")

	users, err := m.userService.ReadAll()
	if err != nil {
		return err
	}

	for _, user := range users {
		currentAuthorizations := authorization.DefaultPortainerAuthorizations()
		currentAuthorizations[portainer.OperationPortainerUserListToken] = true
		currentAuthorizations[portainer.OperationPortainerUserCreateToken] = true
		currentAuthorizations[portainer.OperationPortainerUserRevokeToken] = true
		user.PortainerAuthorizations = currentAuthorizations
		err = m.userService.Update(user.ID, &user)
		if err != nil {
			return err
		}
	}

	return nil
}
