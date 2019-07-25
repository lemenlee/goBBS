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

	rqp := requestParam{uri: tokenURI, body: userLg}
	code, body := rqp.postJSON()

	assert.Equal(t, http.StatusOK, code)
	response := &tokenResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}

	userLg = data.UserLogin{Username: user.Username, PasswordHash: "12345"}
	rqp.body = userLg
	code, body = rqp.postJSON()
	assert.NotEqual(t, http.StatusOK, code)
}

func TestUpdateToken(t *testing.T) {
	user := createTestUser()
	token := data.JwtToken{}
	token.GenerateToken(user.UserRegister.UserLogin)
	rqp := requestParam{uri: tokenURI, auth: token.RefreshTokenStr}
	code, body := rqp.putJSON()
	assert.Equal(t, http.StatusOK, code)
	response := &token.AccessToken
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}

	rqp.auth = token.AccessTokenStr
	code, body = rqp.putJSON()
	assert.Equal(t, http.StatusUnauthorized, code)
}
