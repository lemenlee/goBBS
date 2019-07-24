package api

import (
	"bbs/data"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//UserRoutes 用户路由
func UserRoutes(route *gin.RouterGroup) {
	userRoute := route.Group("/user")
	{
		userRoute.POST("/", createUser)
		userRoute.Use(jwtTokenAuth())
		userRoute.PUT("/", updateUser)
		userRoute.GET("/:username", getUser)
	}
}

//CreateUser 注册用户
func createUser(c *gin.Context) {
	// println("test createuser")
	user := data.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		ErrorsForbidden(c)
	} else {
		_, err := data.FindUser(user.Username)
		if err != nil {
			user.GeneratePasswordHash()
			user.CreateUserDB()
			c.JSON(http.StatusOK, user)
		} else {
			ErrorsUserExist(c)
		}
	}
}

func getUser(c *gin.Context) {
	user, err := data.FindUser(c.Param("username"))
	if err != nil {
		fmt.Println(err)
		ErrorsPageNotFound(c)
		return
	}
	err = c.ShouldBindJSON(&user)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	user.GetUserInfo()
	c.JSON(http.StatusOK, user.UserInfo)
}

func updateUser(c *gin.Context) {
	user := data.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		ErrorsForbidden(c)
	} else {
		user.UpdateUserInfo()
		c.JSON(http.StatusOK, user)
	}
}
