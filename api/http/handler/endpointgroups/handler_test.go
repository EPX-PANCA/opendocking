package endpointgroups

import (
	"testing"

	"github.com/opendocking/opendocking/api/datastore"
	"github.com/opendocking/opendocking/api/internal/testhelpers"
)

func setUpHandler(t *testing.T, store *datastore.Store) *Handler {
	t.Helper()
	handler := NewHandler(testhelpers.NewTestRequestBouncer())
	handler.DataStore = store
	return handler
}
