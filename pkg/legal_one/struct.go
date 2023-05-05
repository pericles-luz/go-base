package legal_one

import "time"

type AuthResponse struct {
	RefreshTokenExpiresIn string   `json:"refresh_token_expires_in,omitempty"`
	APIProductList        string   `json:"api_product_list,omitempty"`
	APIProductListJSON    []string `json:"api_product_list_json,omitempty"`
	OrganizationName      string   `json:"organization_name,omitempty"`
	DeveloperEmail        string   `json:"developer.email,omitempty"`
	TokenType             string   `json:"token_type,omitempty"`
	IssuedAt              string   `json:"issued_at,omitempty"`
	ClientID              string   `json:"client_id,omitempty"`
	AccessToken           string   `json:"access_token,omitempty"`
	ApplicationName       string   `json:"application_name,omitempty"`
	Scope                 string   `json:"scope,omitempty"`
	ExpiresIn             string   `json:"expires_in,omitempty"`
	RefreshCount          string   `json:"refresh_count,omitempty"`
	Status                string   `json:"status,omitempty"`
}

type ContactResponse struct {
	Value []Contact `json:"value,omitempty"`
}

type Contact struct {
	Type                 string    `json:"type,omitempty"`
	ID                   int       `json:"id,omitempty"`
	Name                 string    `json:"name,omitempty"`
	CreationDate         time.Time `json:"creationDate,omitempty"`
	IdentificationNumber string    `json:"identificationNumber,omitempty"`
	Reason               any       `json:"reason,omitempty"`
	ExternalCode         string    `json:"externalCode,omitempty"`
	Notes                any       `json:"notes,omitempty"`
	RexMonitoring        bool      `json:"rexMonitoring,omitempty"`
	CountryID            int       `json:"countryId,omitempty"`
	Queries              []query   `json:"queries,omitempty"`
}

type query struct {
	IsActive    bool `json:"isActive,omitempty"`
	QueryString any  `json:"queryString,omitempty"`
}

type Individual struct {
	BirthDate                       any    `json:"birthDate,omitempty"`
	Gender                          any    `json:"gender,omitempty"`
	Nacionality                     any    `json:"nacionality,omitempty"`
	NitPisPasep                     any    `json:"nitPisPasep,omitempty"`
	PersonStateIdentificationNumber any    `json:"personStateIdentificationNumber,omitempty"`
	VoterID                         any    `json:"voterId,omitempty"`
	VoterGeoraphicZone              any    `json:"voterGeoraphicZone,omitempty"`
	VoterSection                    any    `json:"voterSection,omitempty"`
	WorkerNumber                    any    `json:"workerNumber,omitempty"`
	WorkerSeries                    any    `json:"workerSeries,omitempty"`
	ID                              int    `json:"id,omitempty"`
	Name                            string `json:"name,omitempty"`
	CreationDate                    string `json:"creationDate,omitempty"`
	IdentificationNumber            string `json:"identificationNumber,omitempty"`
	Reason                          any    `json:"reason,omitempty"`
	ExternalCode                    any    `json:"externalCode,omitempty"`
	Notes                           any    `json:"notes,omitempty"`
	RexMonitoring                   bool   `json:"rexMonitoring,omitempty"`
	CountryID                       int    `json:"countryId,omitempty"`
}

type ResponseError struct {
	Error struct {
		Code    string        `json:"code,omitempty"`
		Message string        `json:"message,omitempty"`
		Target  string        `json:"target,omitempty"`
		Details []ErrorDetail `json:"details,omitempty"`
	} `json:"error,omitempty"`
}

type ErrorDetail struct {
	Code    string `json:"code,omitempty"`
	Target  string `json:"target,omitempty"`
	Message string `json:"message,omitempty"`
}
