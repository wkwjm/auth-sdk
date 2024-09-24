package resource

import (
	"github.com/houg/go-oauth2-resource/errors"
	"net/http"
	"strings"
)

type Resource struct {
	TokenServ *TokenService
}

var Instance Resource

func Init(config Config) {
	Instance = Resource{
		TokenServ: NewTokenServ(config),
	}
}

// BearerAuth parse bearer token
func (s *Resource) BearerAuth(r *http.Request) (string, bool) {
	auth := r.Header.Get("Authorization")
	prefix := "Bearer "
	token := ""

	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	} else {
		token = r.FormValue("access_token")
	}

	return token, token != ""
}

// ValidationBearerToken validation the bearer tokens
func (s *Resource) ValidationBearerToken(r *http.Request) (*AccessToken, error) {
	accessToken, ok := s.BearerAuth(r)
	if !ok {
		return nil, errors.ErrInvalidAccessToken
	}
	return s.TokenServ.ParseAccessToken(accessToken)
}
