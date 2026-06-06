package update

import (
	"time"

	portainer "github.com/opendocking/opendocking/api"
	httperrors "github.com/opendocking/opendocking/api/http/errors"
	"github.com/opendocking/opendocking/pkg/validate"
)

func ValidateAutoUpdateSettings(autoUpdate *portainer.AutoUpdateSettings) error {
	if autoUpdate == nil {
		return nil
	}

	if autoUpdate.Webhook == "" && autoUpdate.Interval == "" {
		return httperrors.NewInvalidPayloadError("Webhook or Interval must be provided")
	}

	if autoUpdate.Webhook != "" && !validate.IsUUID(autoUpdate.Webhook) {
		return httperrors.NewInvalidPayloadError("invalid Webhook format")
	}

	if autoUpdate.Interval == "" {
		return nil
	}

	if d, err := time.ParseDuration(autoUpdate.Interval); err != nil {
		return httperrors.NewInvalidPayloadError("invalid Interval format")
	} else if d < time.Minute {
		return httperrors.NewInvalidPayloadError("interval must be at least 1 minute")
	}

	return nil
}
