package resource

import (
	"encoding/json"
	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/houg/go-oauth2-resource/pkg/netutil"
	"io"
	"log"
)

type TokenService struct {
	KeyFunc jwt.Keyfunc
}

func (serv *TokenService) ParseAccessToken(tokenString string) (*AccessToken, error) {
	token, err := jwt.Parse(tokenString, serv.KeyFunc)
	if err != nil {
		log.Printf("validate: %s", err)
		return nil, err
	}
	log.Printf("token.Claims: %s \n", token.Claims)
	return &AccessToken{
		Raw:       token.Raw,
		Header:    token.Header,
		Sub:       parseString(token.Claims, "sub"),
		Scopes:    parseStrings(token.Claims, "scope"),
		GrantType: parseString(token.Claims, "authorizationGrantType"),
		ClientId:  parseString(token.Claims, "clientId"),
	}, nil
}

type Config struct {
	JwkUri string `json:"jwk_uri,omitempty"`
	Cache  bool   `json:"cache,omitempty"`
}

func NewConfig(jwkUri string, cache bool) Config {
	return Config{
		JwkUri: jwkUri,
		Cache:  cache,
	}
}

func NewTokenServ(config Config) *TokenService {
	if config.Cache {
		jwkJson, err := requestJwkJson(config.JwkUri)
		if nil == err {
			jwks, err := keyfunc.NewJWKSetJSON(json.RawMessage(jwkJson))
			if err != nil {
				log.Fatalf("Failed to create a keyfunc.Keyfunc.Error: %s", err)
				return nil
			} else {
				return &TokenService{KeyFunc: jwks.Keyfunc}
			}
		}
	}
	jwks, err := keyfunc.NewDefault([]string{config.JwkUri})
	if err != nil {
		log.Fatalf("Failed to create a keyfunc.Keyfunc.Error: %s", err)
		return nil
	} else {
		return &TokenService{KeyFunc: jwks.Keyfunc}
	}
}

func requestJwkJson(jwkUri string) (string, error) {
	header := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := netutil.HttpGet(jwkUri, header)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	if resp.StatusCode == 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		return string(body[:]), nil
	}
	return "", err
}

func parseStrings(claims jwt.Claims, key string) []string {
	var cs []string
	mapClaims := claims.(jwt.MapClaims)
	switch v := mapClaims[key].(type) {
	case string:
		cs = append(cs, v)
	case []string:
		cs = v
	case []interface{}:
		for _, a := range v {
			vs, ok := a.(string)
			if !ok {
				continue
			}
			cs = append(cs, vs)
		}
	}

	return cs
}

func parseString(claims jwt.Claims, key string) string {
	mapClaims := claims.(jwt.MapClaims)
	var (
		ok  bool
		raw interface{}
		iss string
	)
	raw, ok = mapClaims[key]
	if !ok {
		return ""
	}

	iss, ok = raw.(string)
	if !ok {
		return ""
	}
	return iss
}
