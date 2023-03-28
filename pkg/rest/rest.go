package rest

// Import resty into your code and refer it as `resty`.
import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type Rest struct {
	http   *resty.Client
	token  IToken
	config map[string]interface{}
}

type IRest interface {
	Post(payload map[string]interface{}, link string) (*Response, error)
	PostWithContext(payload map[string]interface{}, link string, ctx context.Context) (*Response, error)
	Get(payload map[string]interface{}, link string) (*Response, error)
	PostWithHeader(payload map[string]interface{}, link string, header map[string]string) (*Response, error)
	GetWithHeader(payload map[string]interface{}, link string, header map[string]string) (*Response, error)
	SetToken(token IToken) error
	SetConfig(key string, value string)
	GetConfig(key string) string
	GetConfigData() map[string]string
	getHttp() (*resty.Client, error)
	getToken() (IToken, error)
	// generateToken() (IToken, error)
}

type Response struct {
	code int
	raw  string
}

type IResponse interface {
	GetRaw() string
	GetCode() int
}

func (b *Rest) getHttp() *resty.Client {
	return b.http
}

func (b *Rest) getToken() (IToken, error) {
	if b.token == nil {
		return nil, errors.New("sem token de autenticação")
	}
	if !b.token.IsValid() {
		b.token = nil
		return nil, errors.New("token inválido")
	}
	return b.token, nil
}

func (b *Rest) SetToken(token IToken) error {
	if !token.IsValid() {
		return errors.New("token já está inválido")
	}
	b.token = token
	return nil
}

func (b *Rest) SetConfig(key string, value string) {
	b.config[key] = value
}

func (b *Rest) GetConfig(key string) string {
	return b.config[key].(string)
}

func (b *Rest) GetConfigData() map[string]interface{} {
	return b.config
}

func (b *Rest) Post(payload map[string]interface{}, link string) (*Response, error) {
	token, err := b.getToken()
	if err != nil {
		return nil, err
	}
	resp, err := b.getHttp().R().SetBody(payload).SetAuthToken(token.GetKey()).Post(link)
	if err != nil {
		return nil, err
	}
	fmt.Println("tempo de resposta", resp.Time())
	return &Response{
		code: resp.StatusCode(),
		raw:  resp.String(),
	}, nil
}

func (b *Rest) PostWithContext(payload map[string]interface{}, link string, ctx context.Context) (*Response, error) {
	token, err := b.getToken()
	if err != nil {
		return nil, err
	}
	resp, err := b.getHttp().R().SetContext(ctx).SetBody(payload).SetAuthToken(token.GetKey()).Post(link)
	if err != nil {
		return nil, err
	}
	fmt.Println("tempo de resposta", resp.Time())
	return &Response{
		code: resp.StatusCode(),
		raw:  resp.String(),
	}, nil
}

func (b *Rest) PostWithHeader(payload map[string]interface{}, link string, header map[string]string) (*Response, error) {
	token, err := b.getToken()
	if err != nil {
		return nil, err
	}
	resp, err := b.getHttp().R().SetBody(payload).SetHeaders(header).SetAuthToken(token.GetKey()).Post(link)
	if err != nil {
		return nil, err
	}
	resp.Time()
	return &Response{
		code: resp.StatusCode(),
		raw:  resp.String(),
	}, nil
}

func (b *Rest) PostWithHeaderNoAuth(payload map[string]interface{}, link string, header map[string]string) (*Response, error) {
	resp, err := b.getHttp().R().SetBody(payload).SetHeaders(header).Post(link)
	if err != nil {
		return nil, err
	}
	resp.Time()
	return &Response{
		code: resp.StatusCode(),
		raw:  resp.String(),
	}, nil
}

func (b *Rest) Get(payload map[string]interface{}, link string) (*Response, error) {
	token, err := b.getToken()
	if err != nil {
		return nil, err
	}
	dados := map[string]string{}
	for k, v := range payload {
		switch t := v.(type) {
		case string:
			dados[k] = v.(string)
		case bool:
			if v.(bool) {
				dados[k] = "true"
				continue
			}
			dados[k] = "false"
		default:
			dados[k] = fmt.Sprintf("%v", t)
		}
	}
	resp, err := b.getHttp().R().SetQueryParams(dados).SetAuthToken(token.GetKey()).Get(link)
	if err != nil {
		return nil, err
	}
	resp.Time()
	return &Response{
		code: resp.StatusCode(),
		raw:  resp.String(),
	}, nil
}

func (b *Rest) GetWithHeader(payload map[string]interface{}, link string, header map[string]string) (*Response, error) {
	token, err := b.getToken()
	if err != nil {
		return nil, err
	}
	dados := map[string]string{}
	for k, v := range payload {
		switch t := v.(type) {
		case string:
			dados[k] = v.(string)
		case bool:
			if v.(bool) {
				dados[k] = "true"
				continue
			}
			dados[k] = "false"
		default:
			dados[k] = fmt.Sprintf("%v", t)
		}
	}
	resp, err := b.getHttp().R().SetQueryParams(dados).SetHeaders(header).SetAuthToken(token.GetKey()).Get(link)
	if err != nil {
		return nil, err
	}
	resp.Time()
	return &Response{
		code: resp.StatusCode(),
		raw:  resp.String(),
	}, nil
}

func (b *Rest) GetWithHeaderNoAuth(payload map[string]interface{}, link string, header map[string]string) (*Response, error) {
	dados := map[string]string{}
	for k, v := range payload {
		switch t := v.(type) {
		case string:
			dados[k] = v.(string)
		case bool:
			if v.(bool) {
				dados[k] = "true"
				continue
			}
			dados[k] = "false"
		default:
			dados[k] = fmt.Sprintf("%v", t)
		}
	}
	resp, err := b.getHttp().R().SetQueryParams(dados).SetHeaders(header).Get(link)
	if err != nil {
		return nil, err
	}
	resp.Time()
	return &Response{
		code: resp.StatusCode(),
		raw:  resp.String(),
	}, nil
}

func (r *Response) GetRaw() string {
	return r.raw
}

func (r *Response) GetCode() int {
	return r.code
}

// func (b *Rest) generateToken(jwt string, validity int) (IToken, error) {
// 	token := NewToken()
// 	token.SetKey(jwt)
// 	token.SetValidity(time.Now().Local().Add(time.Minute * time.Duration(validity)).Format("2006-01-02 15:04:05"))
// 	b.SetToken(token)
// 	return token, nil
// }

func NewRest(config map[string]interface{}) *Rest {
	client := resty.New()
	if config["InsecureSkipVerify"] != nil && config["InsecureSkipVerify"].(bool) {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: config["InsecureSkipVerify"].(bool)})
	}
	rest := &Rest{
		http:   client,
		config: config,
		token:  &Token{},
	}
	rest.http.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})
	rest.getHttp().SetTimeout(1 * time.Minute)
	return rest
}
