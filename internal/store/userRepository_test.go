package store_test

import (
	"testing"

	"github.com/goocarry/cnord-test/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_SaveUser(t *testing.T) {
	s, teardown := store.TestStore(t, dsn)
	defer teardown("users")

	o, err := s.User().Create( "testuser", "testuserov")
	assert.NoError(t, err)
	assert.NotNil(t, o)
}
