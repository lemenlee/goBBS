package api

import (
	"bbs/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.RouterGroup) {
	commentRoute := router.Group("/comment")
	{
		commentRoute.Use(jwtTokenAuth())
		commentRoute.POST("/", createCommentAPI)
		commentRoute.GET("/", getPostCommentAPI)
		commentRoute.GET("/:commentid", getCommentByIdAPI)
		commentRoute.PUT("/", updateCommentAPI)
		commentRoute.DELETE("/:commentid", deleteCommentAPI)
	}

}

func createCommentAPI(c *gin.Context) {
	comment := data.Comment{}
	err := c.ShouldBind(&comment)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	user, err := data.FindUser(crusername)
	if err != nil {
		ErrorsInvalidUser(c)
		return
	}
	post, err := data.FindPostByID(comment.PostID)
	if err != nil {
		ErrorsInvalidToken(c)
		return
	}

	err = comment.CreateComment(user, post)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	c.JSON(http.StatusOK, comment)
}

func getPostCommentAPI(c *gin.Context) {
	post := data.Post{}
	err := c.ShouldBind(&post)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	post, err = data.FindPostByID(post.ID)
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

func updateCommentAPI(c *gin.Context) {
	comment := data.Comment{}
	err := c.ShouldBind(&comment)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	err = comment.UpdateComment()
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	c.Status(http.StatusOK)
}

func deleteCommentAPI(c *gin.Context) {
	commentID := c.Param("commentid")
	comment, err := data.GetCommentByID(commentID)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	err = comment.DeleteComment()
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	c.Status(http.StatusOK)
}

func getCommentByIdAPI(c *gin.Context) {
	commentID := c.Param("commentid")
	comment, err := data.GetCommentByID(commentID)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	c.JSON(http.StatusOK, comment)
}
