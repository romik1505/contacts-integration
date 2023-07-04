package contact

import (
	"os"
	"testing"
	_ "week3_docker/internal/config"
)

func MainTest(m *testing.M) {
	os.Exit(m.Run())
}
