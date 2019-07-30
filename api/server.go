package api

import (
	"bbs/data"
	"fmt"

	"github.com/gin-gonic/gin"
)

var crusername string

//SetupRouter 初始化router
func SetupRouter() *gin.Engine {
	gin.DisableConsoleColor()
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	// r.NoRoute(ErrorsPageNotFound)
	apiRoute := r.Group("/api")
	{
		v1 := apiRoute.Group("/v1")
		{
			TokenRoutes(v1)
			UserRoutes(v1)
			PostRoutes(v1)
			CommentRoutes(v1)
		}
	}
	return r
}

func jwtTokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		fmt.Printf("token: %s", tokenStr)
		accessToken := data.AccessToken{AccessTokenStr: tokenStr}
		crusername = accessToken.ValidToken()
		if len(crusername) == 0 {
			ErrorsInvalidToken(c)
		}
	}
}
