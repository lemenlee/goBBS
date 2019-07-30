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
		userRoute.GET("/:username/comments", getUserCommentAPI)
	}
}

// 注册接口 godoc
// @Summary 注册接口，用于用户注册
// @Description 通过email,username和password，注册用户
// @Tags user
// @Accept json
// @Produce json
// @Param userRegister body data.UserRegister true "User Register"
// @Success 200 {object} data.User
// @Router /user [post]
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

// 用户信息接口 godoc
// @Summary 用户信息接口id
// @Security JWTAuth
// @Description 通过username查询用户
// @Tags user
// @Produce json
// @Param username path string true "username"
// @Success 200 {object} data.User
// @Router /user/{username} [get]
func getUser(c *gin.Context) {
	user, err := data.FindUser(c.Param("username"))
	if err != nil {
		fmt.Println(err)
		return
	}
	user.GetUserInfo()
	c.JSON(http.StatusOK, user.UserInfo)
}

// 更新用户信息 godoc
// @Summary 用户信息更新接口
// @Description 通过user模型，更新用户数据
// @Tags user
// @Security JWTAuth
// @Accept json
// @Produce json
// @Param user body data.User true "User"
// @Success 200 {object} data.User
// @Router /user [put]
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

// 获取用户评论 godoc
// @Summary 获取用户评论接口
// @Description 通过username获取评论
// @Tags comment
// @Security JWTAuth
// @Produce json
// @Param username path string true "username"
// @Success 200 {array} data.Comment
// @Router /user/{username}/comments [get]
func getUserCommentAPI(c *gin.Context) {
	user, err := data.FindUser(crusername)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	comments, err := data.GetAllUserComment(user)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	c.JSON(http.StatusOK, comments)
}
