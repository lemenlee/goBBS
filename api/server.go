package api

import (
	"bbs/data"

	"github.com/gin-gonic/gin"
)

//SetupRouter 初始化router
func SetupRouter() *gin.Engine {
	gin.DisableConsoleColor()
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	r.NoRoute(ErrorsPageNotFound)
	apiRoute := r.Group("/api")
	{
		v1 := apiRoute.Group("/v1")
		{
			TokenRoutes(v1)
			UserRoutes(v1)
		}
	}
	return r
}

func jwtTokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		accessToken := data.AccessToken{AccessTokenStr: tokenStr}
		username := accessToken.ValidToken()
		if len(username) == 0 {
			ErrorsInvalidToken(c)
		}
	}
}
