package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ErrorsPageNotFound 404 json
func ErrorsPageNotFound(c *gin.Context) {
	respondWithError(c, http.StatusNotFound, "uri not found")
}

func ErrorsInvalidUser(c *gin.Context) {
	respondWithError(c, http.StatusUnauthorized, "username or password error")
}

//Forbidden json
func ErrorsForbidden(c *gin.Context) {
	respondWithError(c, http.StatusForbidden, "forbidden")
}

func ErrorsUserExist(c *gin.Context) {
	respondWithError(c, http.StatusForbidden, "user is exist")
}

func ErrorsInvalidToken(c *gin.Context) {
	// fmt.Println("=-==-=-=-==-=-=-=-=-=-=-=-")
	respondWithError(c, http.StatusUnauthorized, "Invalid API token")
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
