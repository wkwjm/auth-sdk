# go-oauth2-resource

1、下载
```
go get -u github.com/houg/go-oauth2-resource
```

2、初始化
```
jwksURL := "https://impre.zdxlz.com/seal/oauth2/jwks"
config := resource.NewConfig(jwksURL, true)
resource.Init(config)
```
3、路由集成中间件
```
GrantType: authorization_code client_credentials

route.GET("/admin", middleware.Oauth2ResourceMiddleware([]string{"openid","profile"}, []string{}), func(c *gin.Context) {  
    c.JSON(200, gin.H{"message": "只有管理员可以访问"})  
})
```