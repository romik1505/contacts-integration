package contact

import (
	"os"
	"testing"
	"week3_docker/internal/config"
)

func MainTest(m *testing.M) {
	config.NewConfig()
	os.Exit(m.Run())
}
