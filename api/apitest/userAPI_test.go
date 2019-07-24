package apitest

import (
	"bbs/data"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserURI = "/api/v1/user/"

func TestCreateUser(t *testing.T) {
	data.Db.Delete(&data.User{})
	userRG := data.UserRegister{Email: "lemon@qq.com",
		UserLogin: data.UserLogin{Username: "lemonlee", PasswordHash: "123"}}
	code, _ := postJSON(UserURI, userRG)
	assert.Equal(t, http.StatusOK, code)

	userRG.PasswordHash = ""
	code, _ = postJSON(UserURI, userRG)
	assert.Equal(t, http.StatusForbidden, code)

	userRG.Username = "12"
	userRG.PasswordHash = "lemonlee"
	code, _ = postJSON(UserURI, userRG)
	assert.Equal(t, http.StatusForbidden, code)

	userRG.Username = ""
	userRG.PasswordHash = ""
	code, _ = postJSON(UserURI, userRG)
	assert.Equal(t, http.StatusForbidden, code)
}
