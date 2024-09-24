package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/houg/go-oauth2-resource/pkg/common"
	"github.com/houg/go-oauth2-resource/resource"
	"net/http"
)

func Oauth2ResourceMiddleware(scopes []string, grantTypes []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if scopes == nil || len(scopes) == 0 {
			c.IndentedJSON(http.StatusUnauthorized, common.Result{
				Code: 401,
				Msg:  common.UNAUTHORIZED_ACCESS,
			})
			return
		}
		accessToken, err := resource.Instance.ValidationBearerToken(c.Request)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, common.Result{
				Code: 401,
				Msg:  err.Error(),
			})
			return
		}
		if !accessToken.HasScopes(scopes...) {
			c.IndentedJSON(http.StatusUnauthorized, common.Result{
				Code: 401,
				Msg:  common.UNAUTHORIZED_ACCESS,
			})
			return
		}
		if grantTypes != nil && len(grantTypes) > 0 && !accessToken.HasGrantType(grantTypes...) {
			c.IndentedJSON(http.StatusUnauthorized, common.Result{
				Code: 401,
				Msg:  common.UNAUTHORIZED_ACCESS,
			})
			return
		}
		c.Set("accessToken", accessToken)
		c.Next()
	}
}
