package api

import (
	"bbs/data"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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
		v1.GET("/cpp", download())
	}
	return r
}

func download() gin.HandlerFunc {
	return func(c *gin.Context) {

		f, err := os.Open("cpp.pdf")
		if err != nil {
			fmt.Println("read file fail", err)
			return
		}
		defer f.Close()

		fd, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println("read to fd fail", err)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
		c.Header("Content-Type", "application/text/plain")
		c.Header("Content-Disposition", "attachment; filename=cpp.pdf")
		c.Header("Accept-Length", fmt.Sprintf("%d", len(fd)))
		c.Writer.Write([]byte(fd))
	}
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
