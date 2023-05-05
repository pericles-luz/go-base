package utils_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestFomartCPF(t *testing.T) {
	cpf, err := utils.FormatCPF("00000000191")
	require.NoError(t, err)
	require.Equal(t, "000.000.001-91", cpf)
}

func TestFormatCPFMustFailIfNumberIsNotValid(t *testing.T) {
	_, err := utils.FormatCPF("00000000192")
	require.Error(t, err)
}
