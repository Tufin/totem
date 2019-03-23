package common_test

import (
	"os"
	"testing"

	"github.com/tufin/totem/common"

	"github.com/stretchr/testify/require"
)

const (
	key   = "TEST_CHECK"
	value = "ME"
)

func TestGetEnvOrExit(t *testing.T) {

	require.NoError(t, os.Setenv(key, value))
	require.Equal(t, value, common.GetEnvOrExit(key))
}

func TestGetEnv(t *testing.T) {

	require.NoError(t, os.Setenv(key, value))
	require.Equal(t, value, common.GetEnv(key))
}
