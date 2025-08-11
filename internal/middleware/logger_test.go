package middleware

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMiddlewareLogger(t *testing.T) {
	middlewareLogger := FormatLogger()
	require.NotNil(t, middlewareLogger)
}
