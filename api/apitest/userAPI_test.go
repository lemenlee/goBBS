package apitest

import (
	"bbs/data"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const userURI = "/api/v1/user/"

func TestCreateUser(t *testing.T) {
	data.Db.Delete(&data.User{})
	userRG := data.UserRegister{Email: "lemon@qq.com",
		UserLogin: data.UserLogin{Username: "lemonlee", PasswordHash: "123"}}
	rqParm := requestParam{uri: userURI, body: userRG}
	code, _ := rqParm.postJSON()
	assert.Equal(t, http.StatusOK, code)

	data.Db.Delete(&data.User{})
	userRG.PasswordHash = ""
	rqParm.body = userRG
	// fmt.Printf("-=-==-=-=-=-=-===-=-, %s", rqParm.body)
	code, _ = rqParm.postJSON()
	assert.Equal(t, http.StatusForbidden, code)

	data.Db.Delete(&data.User{})
	userRG.Username = "12"
	userRG.PasswordHash = "lemonlee"
	rqParm.body = userRG
	code, _ = rqParm.postJSON()
	assert.Equal(t, http.StatusForbidden, code)

	data.Db.Delete(&data.User{})
	userRG.Username = ""
	userRG.PasswordHash = ""
	rqParm.body = userRG
	code, _ = rqParm.postJSON()
	assert.Equal(t, http.StatusForbidden, code)
}

func TestGetUser(t *testing.T) {
	user := createTestUser()
	token := data.JwtToken{}
	token.GenerateToken(user.UserLogin)
	rqparm := requestParam{uri: userURI + user.Username, auth: token.AccessTokenStr}
	code, body := rqparm.getJSON()
	assert.Equal(t, http.StatusOK, code)
	response := &data.UserInfo{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}
	// fmt.Printf("response:%v\n", string(body))
}

func TestUpdateUser(t *testing.T) {
	user := createTestUser()
	token := data.JwtToken{}
	token.GenerateToken(user.UserLogin)
	user.Name = "lemon"
	rqparm := requestParam{uri: userURI, auth: token.AccessTokenStr, body: user}
	code, body := rqparm.putJSON()
	assert.Equal(t, http.StatusOK, code)
	response := &data.User{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}
	fmt.Printf("response:%v\n", string(body))
}

func TestGetUserComment(t *testing.T) {
	// user := createTestUser()
	user, token := createTestUserAndToken()
	post := data.CreateTestPost(user)
	post.GeneratePost(user)
	for i := 0; i < 10; i++ {
		comment := data.CreateTestComment()
		comment.CreateComment(user, post)
	}

	rqparm := requestParam{uri: userURI + user.ID + "/comments", auth: token}
	code, body := rqparm.getJSON()
	assert.Equal(t, http.StatusOK, code)
	response := &[]data.Comment{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Error("解析错误")
	}
	if len(*response) != 10 {
		t.Error("comments nums not true")
	}
}
