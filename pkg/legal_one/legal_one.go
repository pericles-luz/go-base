package legal_one

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/pericles-luz/go-base/pkg/rest"
	"github.com/pericles-luz/go-base/pkg/utils"
)

const (
	TOKEN_VALIDITY = 10
)

type LegalOne struct {
	token  *rest.Token
	rest   *rest.Rest
	parser *Parser
}

func (l *LegalOne) Autenticate() error {
	if l.token != nil && l.token.IsValid() {
		return nil
	}
	authPreBase64 := l.getRest().GetConfig("DE_User") + ":" + l.getRest().GetConfig("PW_Senha")
	authBase64 := base64.StdEncoding.EncodeToString([]byte(authPreBase64))
	resp, err := l.getRest().PostWithHeaderNoAuth(nil, l.getRest().GetConfig("LN_Auth"), map[string]string{
		"Authorization": "Basic " + authBase64,
	})
	if err != nil {
		return err
	}
	response, err := l.getParser().AuthResponse(resp.GetRaw())
	if err != nil {
		return err
	}
	token := rest.NewToken()
	token.SetKey(response.AccessToken)
	token.SetValidity(time.Now().UTC().Add(time.Minute * TOKEN_VALIDITY).Format("2006-01-02 15:04:05"))
	l.token = token
	return nil
}

func (l *LegalOne) GetContactByCPF(cpf string) (*ContactResponse, error) {
	cpf, err := utils.FormatCPF(cpf)
	if err != nil {
		return nil, err
	}
	resp, err := l.get(l.getRest().GetConfig("LN_API")+"/contacts?$filter=identificationNumber eq '"+cpf+"'", nil)
	if err != nil {
		return nil, err
	}
	return l.getParser().GetContactResponse(resp.GetRaw())
}

func (l *LegalOne) IndividualRegistrate(data map[string]interface{}) (*Individual, error) {
	l.getParser().setData(data)
	individual, err := l.getParser().IndividualRegistrateRequest()
	if err != nil {
		return nil, err
	}
	send, err := utils.StructToMapInterface(individual)
	if err != nil {
		return nil, err
	}
	resp, err := l.post(l.getRest().GetConfig("LN_API")+"/individuals", send)
	if err != nil {
		return nil, err
	}
	response, err := l.getParser().IndividualRegistrateResponse(resp.GetRaw())
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (l *LegalOne) IndividualDelete(id int) error {
	resp, err := l.delete(l.getRest().GetConfig("LN_API") + "/individuals/" + utils.IntToString(id))
	if err != nil {
		return err
	}
	if resp.GetCode() != http.StatusNoContent {
		return errors.New("error deleting individual")
	}
	return nil
}

func (l *LegalOne) GetLawsuits() (*LawsuitResponse, error) {
	resp, err := l.get(l.getRest().GetConfig("LN_API")+"/lawsuits", nil)
	if err != nil {
		return nil, err
	}
	return l.getParser().GetLawsuitResponse(resp.GetRaw())
}

func (l *LegalOne) GetLawsuitParticipationByContactID(lawsuitID int, contactID int) (*LitigationParticipationResponse, error) {
	resp, err := l.get(l.getRest().GetConfig("LN_API")+"/lawsuits/"+utils.IntToString(lawsuitID)+"/participants/?$filter=contactId eq "+utils.IntToString(contactID), nil)
	if err != nil {
		return nil, err
	}
	return l.getParser().LitigationParticipationResponse(resp.GetRaw())
}

func (l *LegalOne) GetLawsuitByProcessNumber(processNumber string) (*LawsuitResponse, error) {
	resp, err := l.get(l.getRest().GetConfig("LN_API")+"/lawsuits/?$filter=identifierNumber eq '"+processNumber+"'", nil)
	if err != nil {
		return nil, err
	}
	return l.getParser().GetLawsuitResponse(resp.GetRaw())
}

func (l *LegalOne) GetLawsuitByFolder(folder string) (*LawsuitResponse, error) {
	resp, err := l.get(l.getRest().GetConfig("LN_API")+"/lawsuits/?$filter=folder eq '"+folder+"'", nil)
	if err != nil {
		return nil, err
	}
	return l.getParser().GetLawsuitResponse(resp.GetRaw())
}

func (l *LegalOne) GetAppealByFolder(folder string) (*AppealResponse, error) {
	resp, err := l.get(l.getRest().GetConfig("LN_API")+"/appeals/?$filter=folder eq '"+folder+"'", nil)
	if err != nil {
		return nil, err
	}
	return l.getParser().GetAppealResponse(resp.GetRaw())
}

func (l *LegalOne) GetAppealParticipationByContactID(appealID int, contactID int) (*LitigationParticipationResponse, error) {
	resp, err := l.get(l.getRest().GetConfig("LN_API")+"/appeals/"+utils.IntToString(appealID)+"/participants/?$filter=contactId eq "+utils.IntToString(contactID), nil)
	if err != nil {
		return nil, err
	}
	return l.getParser().LitigationParticipationResponse(resp.GetRaw())
}

func (l *LegalOne) GetLitigationByContactID(contactID int) (*LitigationResponse, error) {
	resp, err := l.get(l.getRest().GetConfig("LN_API")+"/litigations?$filter=participants/any(p:p/contactId eq ("+utils.IntToString(contactID)+")  and (p/positionId eq (24) or p/positionId eq (1)))", nil)
	if err != nil {
		return nil, err
	}
	return l.getParser().GetLitigationResponse(resp.GetRaw())
}

func (l *LegalOne) ParticipationRegistrate(data map[string]interface{}) (*LitigationParticipant, error) {
	l.getParser().setData(data)
	participation, err := l.getParser().ParticipationRegistrateRequest()
	if err != nil {
		return nil, err
	}
	send, err := utils.StructToMapInterface(participation)
	if err != nil {
		return nil, err
	}
	resp, err := l.post(l.getRest().GetConfig("LN_API")+"/lawsuits/"+utils.IntToString(l.parser.getData()["ID_Acao"].(int))+"/participants", send)
	if err != nil {
		return nil, err
	}
	response, err := l.getParser().ParticipationRegistrateResponse(resp.GetRaw())
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (l *LegalOne) ParticipationDelete(lawsuitID int, participationID int) error {
	resp, err := l.delete(l.getRest().GetConfig("LN_API") + "/lawsuits/" + utils.IntToString(lawsuitID) + "/participants/" + utils.IntToString(participationID))
	if err != nil {
		return err
	}
	if resp.GetCode() != http.StatusNoContent {
		return errors.New("error deleting participation")
	}
	return nil
}

func (l *LegalOne) AppealParticipationRegistrate(data map[string]interface{}) (*LitigationParticipant, error) {
	l.getParser().setData(data)
	participation, err := l.getParser().ParticipationRegistrateRequest()
	if err != nil {
		return nil, err
	}
	send, err := utils.StructToMapInterface(participation)
	if err != nil {
		return nil, err
	}
	resp, err := l.post(l.getRest().GetConfig("LN_API")+"/appeals/"+utils.IntToString(l.parser.getData()["ID_Acao"].(int))+"/participants", send)
	if err != nil {
		return nil, err
	}
	response, err := l.getParser().ParticipationRegistrateResponse(resp.GetRaw())
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (l *LegalOne) AppealParticipationDelete(appealID int, participationID int) error {
	resp, err := l.delete(l.getRest().GetConfig("LN_API") + "/appeals/" + utils.IntToString(appealID) + "/participants/" + utils.IntToString(participationID))
	if err != nil {
		return err
	}
	if resp.GetCode() != http.StatusNoContent {
		return errors.New("error deleting participation")
	}
	return nil
}

func (l *LegalOne) GetLitigationUpdateByID(lawsuitID int, count int) (*LitigationUpdateResponse, error) {
	resp, err := l.get(l.getRest().GetConfig("LN_API")+"/Updates?$filter=relationships/any(r:r/linkId eq ("+utils.IntToString(lawsuitID)+"))&$orderBy=id desc&$top="+utils.IntToString(count), nil)
	if err != nil {
		return nil, err
	}
	return l.getParser().LitigationUpdateResponse(resp.GetRaw())
}

func (l *LegalOne) getParser() *Parser {
	return l.parser
}

func (l *LegalOne) getRest() *rest.Rest {
	return l.rest
}

func (l *LegalOne) get(url string, data map[string]interface{}) (*rest.Response, error) {
	if err := l.Autenticate(); err != nil {
		return nil, err
	}
	log.Println("data para o GET: ", data)

	l.getRest().SetToken(l.token)
	resp, err := l.getRest().Get(data, url)
	if err != nil {
		return nil, err
	}
	log.Println("resposta do GET: ", resp)
	response, err := l.getParser().ResponseError(resp.GetRaw())
	if err != nil {
		return nil, err
	}
	if response != "" {
		return nil, errors.New(response)
	}
	return resp, nil
}

func (l *LegalOne) post(url string, data map[string]interface{}) (*rest.Response, error) {
	if err := l.Autenticate(); err != nil {
		return nil, err
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	l.getRest().SetToken(l.token)
	log.Println("dataJson para o POST: ", string(dataJson))
	resp, err := l.getRest().Post(data, url)
	if err != nil {
		return nil, err
	}
	log.Println("resposta do POST: ", resp)
	response, err := l.getParser().ResponseError(resp.GetRaw())
	if err != nil {
		return nil, err
	}
	if response != "" {
		return nil, errors.New(response)
	}
	return resp, nil
}

func (l *LegalOne) delete(url string) (*rest.Response, error) {
	if err := l.Autenticate(); err != nil {
		return nil, err
	}
	l.getRest().SetToken(l.token)
	resp, err := l.getRest().Delete(url)
	if err != nil {
		return nil, err
	}
	if resp.GetCode() == http.StatusNoContent {
		return resp, nil
	}
	log.Println("resposta do DELETE: ", resp)
	response, err := l.getParser().ResponseError(resp.GetRaw())
	if err != nil {
		return nil, err
	}
	if response != "" {
		return nil, errors.New(response)
	}
	return resp, nil
}

func NewLegalOne(rest *rest.Rest) *LegalOne {
	return &LegalOne{
		rest:   rest,
		parser: NewParser(),
	}
}
