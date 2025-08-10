package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEnv_ReturnsEnvValue(t *testing.T) {
	t.Run("ReturnsEnvValue", func(t *testing.T) {
		os.Setenv("TEST_ENV_KEY", "test_value")
		defer os.Unsetenv("TEST_ENV_KEY")
		val := GetEnv("TEST_ENV_KEY", "default")
		require.Equal(t, "test_value", val)
	})

	t.Run("ReturnsDefaultValue", func(t *testing.T) {
		os.Unsetenv("TEST_ENV_KEY_NOT_SET")
		val := GetEnv("TEST_ENV_KEY_NOT_SET", "default")
		require.Equal(t, "default", val)
	})
}
