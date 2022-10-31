package store

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/goocarry/cnord-test/internal/config"
)

// TestStore ...
func TestStore(t *testing.T, dsn string) (*Store, func(...string)) {
	t.Helper()

	config := config.GetConfig()
	config.PostgresURL = dsn
	s := New(config, log.Default())
	err := s.Open()
	if err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			_, err := s.db.Exec(fmt.Sprintf(`TRUNCATE "%s" CASCADE`, strings.Join(tables, ", ")))
			if err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
