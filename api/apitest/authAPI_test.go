package apitest

import (
	"bbs/data"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type tokenResponse struct {
	token data.JwtToken
	user  data.UserInfo
}

const tokenURI = "/api/v1/token"

func TestGenerateToken(t *testing.T) {
	user := createTestUser()
	userLg := data.UserLogin{Username: user.Username, PasswordHash: "123456"}
	code, body := postJSON(tokenURI, userLg)

	assert.Equal(t, http.StatusOK, code)
	response := &tokenResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}

	userLg = data.UserLogin{Username: user.Username, PasswordHash: "12345"}
	code, body = postJSON(tokenURI, userLg)
	assert.NotEqual(t, http.StatusOK, code)
}

func TestUpdateToken(t *testing.T) {
	user := createTestUser()
	token := data.JwtToken{}
	token.GenerateToken(user.UserRegister.UserLogin)

	code, body := putHeader(tokenURI, token.RefreshTokenStr)
	assert.Equal(t, http.StatusOK, code)
	response := &token.AccessToken
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}

	code, body = putHeader(tokenURI, token.AccessTokenStr)
	assert.Equal(t, http.StatusUnauthorized, code)
}

func createTestUser() data.User {
	data.Db.Delete(&data.User{})
	user := data.CreateUser("lemon@qq.com", "123456", "lemonlee")
	user.GeneratePasswordHash()
	user.CreateUserDB()
	return user
}
