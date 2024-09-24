package main

import (
	"github.com/houg/go-oauth2-resource/resource"
	"log"
)

func main() {

	// This is a local Keycloak JWK Set endpoint for the master realm.
	jwksURL := "https://impre.zdxlz.com/seal/oauth2/jwks"

	// Get a JWT to parse.
	jwtB64 := "eyJraWQiOiJwRVBhZXZNWHVMNUtxUkMiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiIxNzgxMjUxNzYxMDMzNDE2NzA0IiwiYXVkIjoiOWU5YzI5YzM1NWM2NDZhY2FhZiIsIm5iZiI6MTcyNzE1NjQ2NiwiY2xpZW50SWQiOiI5ZTljMjljMzU1YzY0NmFjYWFmIiwic2NvcGUiOlsib3BlbmlkIiwicHJvZmlsZSJdLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjkwMDAvc2VhbCIsImF1dGhvcml6YXRpb25HcmFudFR5cGUiOiJhdXRob3JpemF0aW9uX2NvZGUiLCJleHAiOjE3MjcxNjM2NjYsImlhdCI6MTcyNzE1NjQ2NiwianRpIjoiMjM3NzMxMDYtYTZjYi00YWRkLWI0ODEtZGE4NzMwOTBjNzM1In0.TSBkcrGe25L9hkGFPZFm9KFQsYH_Ah17kFB2aYscvLl_6vFhbMJzIACl7CRa77Z4-1re9zNqt7hqXuqH1EfutbwzVKBoewhrS3MwnnfEri1lBY7gahfU2UvftW-vr2jbdnIMqH3XkMdRl2kBFgDen9uvH8NBcuwod0LkW2COBG6dQGSfCg6JKDqsZvk_P5rImanw0UqCdQTEpTX_rIWBgzXZWqIzcqlU5jbtKJJs00ahgKu1GE8JzwJPLbULDFLXQQpiTpe7pRpbIgzfym5TYrrnWokziQ30TpIIitjV9bHUjBZqIYRd0-BHFeR5uje6ntAeVgxrFfa1SiQyHchoLQ"

	config := resource.NewConfig(jwksURL, true)
	tokenService := resource.NewTokenServ(config)

	accessToken, err := tokenService.ParseAccessToken(jwtB64)
	if err != nil {
		log.Printf("Error: %s \n", err)
		return
	}
	log.Printf("HasScopes: %v \n", accessToken.HasScopes("openid"))
	log.Printf("HasGrantType: %v \n", accessToken.HasGrantType("authorization_code"))
}
