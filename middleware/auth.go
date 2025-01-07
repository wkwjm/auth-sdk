// @Author wangkang 2025/1/5 14:15:00
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	authTokenHeader = "Authorization"
	proEnv          = "pro"
)

// ErrorResponse 用于封装错误信息，以便返回JSON格式的响应
type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// AuthMiddleware 是实现鉴权逻辑的中间件函数，接收环境变量和鉴权URL作为参数
func AuthMiddleware(environment, url string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录环境和鉴权URL等信息到日志（这里假设使用gin框架自带的日志功能，实际可按需调整日志库）
		log.Printf("environment....%s", environment)
		log.Printf("auth url....%s", url)

		// 非线上环境走sdk鉴权
		if environment != proEnv {
			requestURL := c.Request.URL.String()
			log.Printf("requestURL....%s", requestURL)

			// 检测header中是否携带token
			token := c.Request.Header.Get(authTokenHeader)
			if token != "" {
				// 检测token合法性
				resultMap, err := authenticate(url, authTokenHeader, token)
				if err != nil {
					// 处理验证过程中的错误，比如记录日志等，这里简单返回500错误
					log.Printf("鉴权过程出现错误: %v", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
					c.Abort()
					return
				}
				authenticate, ok := resultMap["authenticate"].(bool)
				if ok && authenticate {
					// 认证通过，获取响应头信息（认证服务器返回的响应头包含要转发的请求头相关配置等）
					customHeaders, ok := resultMap["headers"].(map[string]string)
					if ok {
						for headerName, headerValue := range customHeaders {
							c.Request.Header.Set(headerName, headerValue)
						}
						c.Next()
						return
					}
				}
			}
			// 鉴权失败，返回401 Unauthorized的JSON格式响应
			c.JSON(http.StatusUnauthorized, ErrorResponse{
				Message: "Unauthorized",
				Status:  http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
		// 生产环境直接放行请求
		c.Next()
	}
}

// 认证服务器交互的逻辑
func authenticate(url, headerName, token string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	result["authenticate"] = false
	result["headers"] = []string{}
	client := &http.Client{}
	fullUrl := url
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return result, nil
	}
	// 请求头设置
	req.Header.Set("Accept", "*/*")
	req.Header.Set("connection", "Keep-Alive")
	req.Header.Set("user-agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1;SV1)")
	req.Header.Set("Content-type", "application/json;charset=utf-8")
	req.Header.Set(headerName, token)
	resp, err := client.Do(req)
	if err != nil {
		return result, nil
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, nil
	}
	result["respBody"] = string(respBody)
	if resp.StatusCode == http.StatusOK {
		testMap := make(map[string]string)
		i := 0
		for key, values := range resp.Header {
			for _, value := range values {
				testMap[key] = value
			}
			i++
		}
		fmt.Println(testMap)
		result["headers"] = testMap
		result["authenticate"] = true
	}
	return result, nil
}
