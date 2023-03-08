package utils_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestGetBaseDirectory(t *testing.T) {
	baseDirectory := utils.GetBaseDirectory("config")
	require.NotEmpty(t, baseDirectory)
	require.Contains(t, baseDirectory, "/config")
}
