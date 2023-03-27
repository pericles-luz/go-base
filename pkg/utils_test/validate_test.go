package utils_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestValidateCPF(t *testing.T) {
	require.True(t, utils.ValidateCPF("00000000191"))
	require.False(t, utils.ValidateCPF("1234567890"))
	require.False(t, utils.ValidateCPF("123456789012"))
	require.False(t, utils.ValidateCPF("1234567890a"))
	require.False(t, utils.ValidateCPF("11111111111"))
}

func TestValidateUUID(t *testing.T) {
	require.True(t, utils.ValidateUUID("5d81243d-0e3a-41ac-a250-ed147d24f15d"))
	require.False(t, utils.ValidateUUID("5d812433d-0e3a-41ac-a250-ed147d24f15d"))
	require.False(t, utils.ValidateUUID("5d81243d-0e63a-41ac-a250-ed147d24f15d"))
	require.False(t, utils.ValidateUUID("5d81243d-0e3a-413ac-a250-ed147d24f15d"))
	require.False(t, utils.ValidateUUID("5d81243d-0e3a-41ac-a2503-ed147d24f15d"))
	require.False(t, utils.ValidateUUID("5d81243d-0e3a-41ac-a250-e3d147d24f15d"))
}

func TestValidateTimestamp(t *testing.T) {
	require.True(t, utils.ValidateTimestamp("2020-01-01 00:00:00"))
	require.False(t, utils.ValidateTimestamp("2020-01-01T00:00:00Z"))
	require.False(t, utils.ValidateTimestamp("2020-21-01 00:00:00"))
	require.False(t, utils.ValidateTimestamp("2020-02-30 00:00:00"))
	require.False(t, utils.ValidateTimestamp("2020-01-01T00:00:00"))
	require.False(t, utils.ValidateTimestamp("2020-01-01T00:00:00.000Z"))
	require.False(t, utils.ValidateTimestamp("2020-01-01T00:00:00.000"))
}

func TestValidateDD(t *testing.T) {
	require.True(t, utils.ValidateDDD("11"))
	require.False(t, utils.ValidateDDD("1"))
	require.False(t, utils.ValidateDDD("111"))
	require.False(t, utils.ValidateDDD("a1"))
}

func TestValidatePhoneNumber(t *testing.T) {
	require.True(t, utils.ValidatePhoneNumber("123456789"))
	require.True(t, utils.ValidatePhoneNumber("12345678"))
	require.False(t, utils.ValidatePhoneNumber("1234567890"))
	require.False(t, utils.ValidatePhoneNumber("12345678a"))
	require.False(t, utils.ValidatePhoneNumber("1234567"))
	require.False(t, utils.ValidatePhoneNumber("1234567a"))
	require.False(t, utils.ValidatePhoneNumber("1234567"))
}

func TestValidateEmail(t *testing.T) {
	require.True(t, utils.ValidateEmail("email@test.com"))
	require.False(t, utils.ValidateEmail("emailtest.com"))
	require.False(t, utils.ValidateEmail("email@test"))
	require.False(t, utils.ValidateEmail("ema"))
}

func TestValidateURL(t *testing.T) {
	require.True(t, utils.ValidateURL("http://www.google.com"))
	require.True(t, utils.ValidateURL("https://www.google.com"))
	require.True(t, utils.ValidateURL("https://www.google.com/"))
	require.True(t, utils.ValidateURL("https://www.google.com/test"))
	require.True(t, utils.ValidateURL("https://www.google.com/test/"))
	require.False(t, utils.ValidateURL("https//www.google.com/test/test"))
}
