package utils_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateValidToken(t *testing.T) {
	server := utils.NewJwtServer("secret")
	token, err := server.Create(map[string]interface{}{
		"user_id": "1",
		"sub":     "teste",
	})
	require.NoError(t, err)
	require.True(t, server.Valid(token))
	require.False(t, server.Valid(token[:len(token)-1]))
	parser, err := server.Parse(token)
	require.NoError(t, err)
	require.Equal(t, "1", parser["user_id"])
}

func TestExtractValue(t *testing.T) {
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIwMDAwMDAwMDE5MSIsImlzcyI6InNpbmRpcmVjZWl0YS5vcmcuYnIiLCJkYWRvcyI6eyJjcGYiOiIwMDAwMDAwMDE5MSIsIm5vbWUiOiJKb2FxdWltIGRlIFRlc3RlIiwidW5pZGFkZVNpbmRpY2FsIjp7ImlkIjoiNjUwODMwMjctYzk3OC00OThkLTk3NGItOWVlMDRiYWY3YjM1Iiwibm9tZSI6IkRTIERFIFRFU1RFIn19LCJleHAiOjE1MTYyMzkwMjJ9.pas0t80V6B7Shw_pzvmjBWEzIQYr1leRlu03Kxph2-U"
	pessoa, err := utils.ExtractValue("dados", jwt)
	require.NoError(t, err)
	require.Equal(t, utils.TEST_CPF, pessoa.(map[string]interface{})["cpf"].(string))
	require.Equal(t, "Joaquim de Teste", pessoa.(map[string]interface{})["nome"].(string))
	require.Equal(t, "65083027-c978-498d-974b-9ee04baf7b35", pessoa.(map[string]interface{})["unidadeSindical"].(map[string]interface{})["id"].(string))
	require.Equal(t, "DS DE TESTE", pessoa.(map[string]interface{})["unidadeSindical"].(map[string]interface{})["nome"].(string))
	sub, err := utils.ExtractValue("sub", jwt)
	require.NoError(t, err)
	require.Equal(t, utils.TEST_CPF, sub)
}
