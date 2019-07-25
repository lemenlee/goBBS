package api

import (
	"bbs/data"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//PostRoutes 用户路由
func PostRoutes(route *gin.RouterGroup) {
	postRoute := route.Group("/post")
	{
		postRoute.Use(jwtTokenAuth())
		postRoute.POST("/", createPost)
		postRoute.PUT("/", updatePost)
		postRoute.GET("/", getAllPost)
		postRoute.GET("/:postid", getPost)
		postRoute.DELETE("/:postid", deletePost)
	}
}

func createPost(c *gin.Context) {

	postModel := data.PostModel{}
	err := c.ShouldBindJSON(&postModel)
	post := data.Post{PostModel: postModel}
	if err != nil {

		ErrorsForbidden(c)
		return
	}
	user, err := data.FindUser(crusername)
	if err != nil {
		ErrorsInvalidUser(c)
		return
	}
	err = post.GeneratePost(user)
	if err != nil {
		fmt.Println(err.Error())
		ErrorsForbidden(c)
		return
	}
	c.JSON(http.StatusOK, post)
	// c.Status(http.StatusOK)
}

func updatePost(c *gin.Context) {
	post := data.Post{}
	err := c.ShouldBindJSON(&post)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	err = post.UpdatePost()
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	c.Status(http.StatusOK)
}

func getAllPost(c *gin.Context) {
	user, _ := data.FindUser(crusername)
	posts, err := data.FindUserAllPost(user)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	c.JSON(http.StatusOK, posts)
}

func getPost(c *gin.Context) {
	// user, _ := data.FindUser(crusername)
	postID := c.Param("postid")
	post, err := data.FindPostByID(postID)
	fmt.Println("=-=-=-=-==-=-==-=-=-=-=-=-=-=-=-=-=-=")
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	c.JSON(http.StatusOK, post)
}

func deletePost(c *gin.Context) {
	postID := c.Param("postid")
	post, err := data.FindPostByID(postID)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	err = post.DeletePost()
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	c.Status(http.StatusOK)
}
