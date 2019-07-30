package api

import (
	"bbs/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TokenRoutes Token路由
func TokenRoutes(route *gin.RouterGroup) {
	route.POST("/token", getToken)
	route.PUT("/token", updateAccessToken)
}

// 登录接口 godoc
// @Summary 登录接口，获取令牌和用户信息
// @Description 通过username和password，获取信息
// @Tags auth
// @Accept json
// @Produce json
// @Param userInfo body data.UserLogin true "User account"
// @Success 200 {object} data.UserInfo
// @Router /token [post]
func getToken(c *gin.Context) {
	userLogin := data.UserLogin{}
	err := c.ShouldBindJSON(&userLogin)
	if err != nil {
		// fmt.Println(err.Error())
		ErrorsForbidden(c)
	} else {
		if userLogin.ValidateUser() {
			token := data.JwtToken{}
			err := token.GenerateToken(userLogin)
			if err != nil {
				ErrorsForbidden(c)
			}
			user, _ := data.FindUser(userLogin.Username)
			c.JSON(http.StatusOK, gin.H{
				"user":  user.UserInfo,
				"token": token,
			})
		} else {
			// println("123 not found")
			ErrorsInvalidUser(c)
		}
	}
}

// 更新Token godoc
// @Summary 更新Token接口
// @Description 通过RefreshToken，更新AccessToken
// @Tags auth
// @Security JWTAuth
// @Produce json
// @Success 200
// @Router /token [put]
//UpdateAccessToken 更新token权限
func updateAccessToken(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	refreshToken := data.RefreshToken{RefreshTokenStr: tokenStr}

	username := refreshToken.ValidToken()
	user, err := data.FindUser(username)
	// fmt.Printf("tokenstr: %s, username: %s", tokenStr, user.Username)
	if err != nil {
		ErrorsInvalidToken(c)
		return
	}
	accesToken := data.AccessToken{}

	err = accesToken.GenerateToken(user.UserLogin)
	if err != nil {
		ErrorsForbidden(c)
	}
	c.JSON(http.StatusOK, accesToken)
}
