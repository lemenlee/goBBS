package api

import (
	"bbs/data"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CommentRoutes 评论系统路由
func CommentRoutes(router *gin.RouterGroup) {

	commentRoute := router.Group("/comment")
	{
		commentRoute.Use(jwtTokenAuth())

		commentRoute.POST("/", createCommentAPI)
		commentRoute.GET("/:commentid", getCommentByIDAPI)
		commentRoute.PUT("/", updateCommentAPI)
		commentRoute.DELETE("/:commentid", deleteCommentAPI)
	}

}

// 发布评论 godoc
// @Summary 发布评论接口
// @Description 通过comment模型，发布评论
// @Tags comment
// @Security JWTAuth
// @Accept json
// @Produce json
// @Param post body data.Comment true "Comment"
// @Success 200 {object} data.Comment
// @Router /comment [post]
func createCommentAPI(c *gin.Context) {
	fmt.Println("-=-=-=-=-=-==-----------------")
	comment := data.Comment{}
	err := c.ShouldBindJSON(&comment)
	fmt.Printf("comment: %s", comment)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	user, err := data.FindUser(crusername)

	if err != nil {
		fmt.Printf("user:%s, err:%s", user.Username, err)
		ErrorsInvalidUser(c)
		return
	}
	post, err := data.FindPostByID(comment.PostID)
	if err != nil {
		fmt.Printf("post:%s, err:%s", comment.PostID, err)
		ErrorsInvalidToken(c)
		return
	}
	// comment1 := data.Comment{CommentModel: comment}
	err = comment.CreateComment(user, post)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	c.JSON(http.StatusOK, comment)
}

// 更新评论 godoc
// @Summary 更新评论接口
// @Description 通过comment模型，更新评论
// @Tags comment
// @Security JWTAuth
// @Accept json
// @Produce json
// @Param post body data.Comment true "Comment"
// @Success 200 {object} data.Comment
// @Router /comment [put]
func updateCommentAPI(c *gin.Context) {
	comment := data.Comment{}
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	err = comment.UpdateComment()
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	c.JSON(http.StatusOK, comment)
}

// 删除评论 godoc
// @Summary 发布评论接口
// @Description 通过commentid，删除评论
// @Tags comment
// @Security JWTAuth
// @Accept json
// @Produce json
// @Param commentid path string true "commentid"
// @Success 200
// @Router /comment/{commentid} [delete]
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

// 获取指定评论 godoc
// @Summary 获取指定评论接口
// @Description 通过commentid，获取评论
// @Tags comment
// @Security JWTAuth
// @Accept json
// @Produce json
// @Param commentid path string true "commentid"
// @Success 200 {object} data.Comment
// @Router /comment/{commentid} [get]
func getCommentByIDAPI(c *gin.Context) {
	commentID := c.Param("commentid")
	comment, err := data.GetCommentByID(commentID)
	if err != nil {
		ErrorsForbidden(c)
		return
	}
	c.JSON(http.StatusOK, comment)
}
