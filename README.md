# go-oauth2-resource

1、下载
```
go get -u github.com/wkwjm/auth-sdk
```

2、初始化
```
// 设置环境变量和鉴权URL，这里示例中简单硬编码，实际可从配置文件等地方获取
environment := "pre"
authURL := "https://impre.zdxlz.com/seal/v1/authentication"
log.Printf("start: %v \n", authURL)
```
3、路由集成中间件
```
// 使用中间件，例如对所有的请求都应用鉴权中间件，可根据实际需求调整路由分组等应用方式
r.Use(middleware.AuthMiddleware(environment, authURL))
```
4、示例
```
func main() {
	r := gin.Default()

	// 设置环境变量和鉴权URL，这里示例中简单硬编码，实际可从配置文件等地方获取
	environment := "pre"
	authURL := "https://impre.zdxlz.com/seal/v1/authentication"
	log.Printf("start: %v \n", authURL)

	// 使用中间件，例如对所有的请求都应用鉴权中间件，可根据实际需求调整路由分组等应用方式
	r.Use(middleware.AuthMiddleware(environment, authURL))

	// 定义一个简单的路由处理函数，模拟业务接口
	r.GET("/hello", func(c *gin.Context) {
		app := c.Request.Header.Get("X-App-Id")
		log.Printf(app)
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	r.Run(":8080")
}

```