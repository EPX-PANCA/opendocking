package git

import (
	gittypes "github.com/opendocking/opendocking/api/git/types"
)

func GetCredentials(auth *gittypes.GitAuthentication) (string, string) {
	if auth == nil {
		return "", ""
	}

	return auth.Username, auth.Password
}
