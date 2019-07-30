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
		postRoute.GET("/:postid/comments", getPostCommentAPI)
	}
}

// 发布文章 godoc
// @Summary 发布文章接口
// @Description 通过postModel模型，发布文章
// @Tags post
// @Security JWTAuth
// @Accept json
// @Produce json
// @Param post body data.PostModel true "PostModel"
// @Success 200 {object} data.Post
// @Router /post [post]
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

// 更新文章 godoc
// @Summary 更新文章接口
// @Description 通过post模型，更新文章
// @Tags post
// @Security JWTAuth
// @Accept json
// @Produce json
// @Param post body data.Post true "Post"
// @Success 200
// @Router /post [put]
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

// 获取当前用户所有文章 godoc
// @Summary 发布文章接口
// @Description 获取当前所有文章
// @Tags post
// @Security JWTAuth
// @Produce json
// @Success 200 {array} data.Post
// @Router /post [get]
func getAllPost(c *gin.Context) {
	user, _ := data.FindUser(crusername)
	posts, err := data.FindUserAllPost(user)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	c.JSON(http.StatusOK, posts)
}

// 通过文章ID获取文章 godoc
// @Summary 获取指定文章接口
// @Description 通过postid获取文章
// @Tags post
// @Security JWTAuth
// @Produce json
// @Param postid path string true "postid"
// @Success 200 {object} data.Post
// @Router /post/{postid} [get]
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

// 通过文章ID删除文章 godoc
// @Summary 删除指定文章id
// @Description 通过postid删除文章
// @Tags post
// @Security JWTAuth
// @Produce json
// @Param postid path string true "postid"
// @Success 200 {object} data.Post
// @Router /post/{postid} [delete]
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

// 通过文章ID获取文章 godoc
// @Summary 获取指定文章接口
// @Description 通过postid获取文章
// @Tags post
// @Security JWTAuth
// @Produce json
// @Param postid path string true "postid"
// @Success 200 {array} data.Comment
// @Router /post/{postid}/comments [get]
func getPostCommentAPI(c *gin.Context) {
	postid := c.Param("postid")
	post, err := data.FindPostByID(postid)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	comments, err := data.GetAllPostComment(post)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	c.JSON(http.StatusOK, comments)
}
