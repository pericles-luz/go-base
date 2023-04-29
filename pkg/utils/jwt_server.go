package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtServer struct {
	Secret string
}

type JwtServerInterface interface {
	Valid(token string) bool
	Create(payload map[string]interface{}) (string, error)
	Parse(token string) (map[string]interface{}, error)
}

func NewJwtServer(secret string) *JwtServer {
	return &JwtServer{
		Secret: fmt.Sprintf("%s%d", secret, time.Now().YearDay()),
	}
}

// Validate a token
func (j *JwtServer) Valid(token string) bool {
	_, err := j.Parse(token)
	if err != nil {
		log.Println("invalid token:", err)
	}
	return err == nil
}

// Create a token with a given payload.
func (j *JwtServer) Create(payload map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))
	sub, ok := payload["sub"].(string)
	if !ok {
		return "", errors.New("sub not found")
	}
	return token.SignedString([]byte(Hash256(sub + j.Secret)))
}

// Extract the payload from a token.
// If the token is invalid, return an error.
func (j *JwtServer) Parse(token string) (map[string]interface{}, error) {
	if len(token) == 0 {
		return nil, errors.New("token is empty")
	}
	sub, err := ExtractValue("sub", token)
	if err != nil {
		return nil, err
	}
	claims := jwt.MapClaims{}
	parser, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		hash := Hash256(sub.(string) + j.Secret)
		return []byte(hash), nil
	})
	if err != nil {
		return nil, err
	}
	if !parser.Valid {
		return nil, errors.New("token is invalid")
	}
	if parser.Method.Alg() != jwt.SigningMethodHS256.Alg() {
		return nil, errors.New("token is invalid")
	}
	return claims, nil
}

// Extract a value from a token.
func ExtractValue(key string, jwt string) (interface{}, error) {
	// separate jwt payload
	payload := strings.Split(jwt, ".")[1]
	// decode payload
	decoded, err := base64.RawStdEncoding.DecodeString(payload)
	if err != nil {
		log.Println("extraction incorrect:", err, string(decoded))
		return nil, err
	}
	// convert to map
	var data map[string]interface{}
	err = json.Unmarshal(decoded, &data)
	if err != nil {
		return nil, err
	}
	// return value
	return data[key], nil
}
