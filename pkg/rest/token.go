package rest

import (
	"time"
)

type Token struct {
	validity time.Time
	key      string
}

type IToken interface {
	SetValidity(validity string) error
	GetValidity() string
	IsValid() bool
	GetKey() string
	SetKey(key string)
}

func (t *Token) SetValidity(validity string) error {
	dtValidity, err := time.Parse("2006-01-02 15:04:05", validity)
	if nil == err {
		t.validity = dtValidity
	}
	return err
}

func (t *Token) SetKey(key string) {
	t.key = key
}

func (t *Token) IsValid() bool {
	if len(t.key) == 0 {
		return false
	}
	isValid := time.Now().Before(t.validity.Add(3 * time.Hour))
	return isValid
}

func (t *Token) GetValidity() string {
	return t.validity.Format("2006-01-02 15:04:05")
}

func (t *Token) GetKey() string {
	return t.key
}

func NewToken() *Token {
	return &Token{}
}
