package legal_one_test

import (
	"os"
	"testing"
	"time"

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
	contact, err := legalOne.GetContactByCPF("000.000.001-91")
	require.NoError(t, err)
	require.Equal(t, "000.000.001-91", contact.Value[0].IdentificationNumber)
	t.Log(contact)
}

func TestLegalOneGetLawsuits(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	lawsuits, err := legalOne.GetLawsuits()
	require.NoError(t, err)
	require.NotEmpty(t, lawsuits.Value)
}

func TestLegalOneIndividualRegistrate(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	data := map[string]interface{}{
		"DE_Pessoa": "Joaquim de Teste",
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

func TestLegalOneGetLawsuitParticipationByContactID(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	participations, err := legalOne.GetLawsuitParticipationByContactID(1, 2888)
	require.NoError(t, err)
	require.NotEmpty(t, participations.Value)
	t.Log(participations)
}

func TestLegalOneGetLawsuitByProcessNumber(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	lawsuits, err := legalOne.GetLawsuitByProcessNumber("1004915-65.2018.4.01.3400")
	require.NoError(t, err)
	require.NotEmpty(t, lawsuits.Value)
	t.Log(lawsuits)
}

func TestLegalOneGetLawsuitByFolder(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	lawsuits, err := legalOne.GetLawsuitByFolder("Colet-0295")
	require.NoError(t, err)
	require.NotEmpty(t, lawsuits.Value)
	t.Log(lawsuits)
}

func TestLegalOneParticipationRegistrate(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	data := map[string]interface{}{
		"ID_Contato": 25089,
		"ID_Acao":    1,
	}
	participation, err := legalOne.ParticipationRegistrate(data)
	t.Log(participation)
	require.NoError(t, err)
	require.NotEmpty(t, participation.ID)
	time.Sleep(1 * time.Second)
	require.NoError(t, legalOne.ParticipationDelete(1, participation.ID))
}

func TestLegalOneParticipationDelete(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	legalOne, err := factory.NewLegalOne("legalone.prod")
	require.NoError(t, err)
	require.NoError(t, legalOne.ParticipationDelete(1, 349846))
}
