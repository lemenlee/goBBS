package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//ErrorsPageNotFound 404 json
func ErrorsPageNotFound(c *gin.Context) {
	respondWithError(c, http.StatusNotFound, "uri not found")
}

//ErrorsInvalidUser 非法用户
func ErrorsInvalidUser(c *gin.Context) {
	respondWithError(c, http.StatusUnauthorized, "username or password error")
}

//ErrorsForbidden Forbidden json
func ErrorsForbidden(c *gin.Context) {
	respondWithError(c, http.StatusForbidden, "forbidden")
}

//ErrorsUserExist User is exist
func ErrorsUserExist(c *gin.Context) {
	fmt.Println("=-==-=-=-==-=-=-=-=-=-=-=-")
	respondWithError(c, http.StatusForbidden, "user is exist")
}

//ErrorsInvalidToken 非法token
func ErrorsInvalidToken(c *gin.Context) {
	//
	respondWithError(c, http.StatusUnauthorized, "Invalid API token")
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
