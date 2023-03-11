package serving

import (
	"net/http"
	"strings"

	"github.com/pericles-luz/go-base/pkg/infra/conf"
	"github.com/pericles-luz/go-base/pkg/utils"
)

type JWTMiddleware struct {
	cfg *conf.Config
}

func NewJWTMiddleware(cfg *conf.Config) *JWTMiddleware {
	return &JWTMiddleware{
		cfg: cfg,
	}
}

func (j *JWTMiddleware) Validate(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if j.dismissToken(r) {
		next(w, r)
		return
	}
	authorization := r.Header.Get("Authorization")
	if len(authorization) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !strings.HasPrefix(strings.ToLower(authorization), "bearer ") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	secret := j.cfg.JwtSecret
	if !utils.NewJwtServer(secret).Valid(strings.Split(authorization, " ")[1]) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	next(w, r)
}

func (j *JWTMiddleware) dismissToken(r *http.Request) bool {
	if r.Method == http.MethodOptions {
		return true
	}
	if strings.HasSuffix(r.URL.Path, "/") {
		return true
	}
	noAuth := []string{
		"/autenticar",
		"/autenticar/token",
		"/certificado",
	}
	for _, path := range noAuth {
		if strings.HasSuffix(r.URL.Path, path) {
			return true
		}
	}
	return false
}
