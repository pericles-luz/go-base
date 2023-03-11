package server

import (
	"net/http"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/pericles-luz/go-base/pkg/infra/conf"
	"github.com/rs/cors"
)

func NewCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			if strings.HasPrefix(origin, "https://") && strings.HasSuffix(origin, ".sindireceita.org.br") {
				return true
			}
			return false
		},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		Debug:            true,
	})
}

func NewNeroni(c *cors.Cors, cfg *conf.Config) *negroni.Negroni {
	jwt := NewJWTMiddleware(cfg)
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.Use(c)
	n.Use(negroni.HandlerFunc(jwt.Validate))
	return n
}
