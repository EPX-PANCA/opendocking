package libhelm

import (
	"github.com/opendocking/opendocking/pkg/libhelm/sdk"
	"github.com/opendocking/opendocking/pkg/libhelm/types"
)

// NewHelmPackageManager returns a new instance of HelmPackageManager based on HelmConfig
func NewHelmPackageManager() types.HelmPackageManager {
	return sdk.NewHelmSDKPackageManager()
}
