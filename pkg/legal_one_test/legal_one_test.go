package legal_one_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-base/internals/factory"
	"github.com/stretchr/testify/require"
)

func TestLegalOneAuthentication(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	require.NoError(t, legalOne.Autenticate())
}

func TestLegalOneGetContactByCPF(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	contact, err := legalOne.GetContactByCPF("91134846649")
	require.NoError(t, err)
	require.Equal(t, "911.348.466-49", contact.Value[0].IdentificationNumber)
}

func TestLegalOneIndividualRequstrate(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	data := map[string]interface{}{
		"DE_Pessoa": "Pericles Luz",
		"CO_CPF":    "00000000191",
	}
	individual, err := legalOne.IndividualRegistrate(data)
	require.NoError(t, err)
	require.Equal(t, "000.000.001-91", individual.IdentificationNumber)
	require.Equal(t, "Joaquim de Teste", individual.Name)
	t.Log(individual)
}

func TestLegalOneIndividualDelete(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	require.NoError(t, legalOne.IndividualDelete(25087))
}
