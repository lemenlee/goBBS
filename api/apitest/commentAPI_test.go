package apitest

import (
	"bbs/data"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const commentURI string = "/api/v1/comment/"

func TestCreateCommentAPI(t *testing.T) {
	user, token := createTestUserAndToken()
	post := data.CreateTestPost(user)
	post.GeneratePost(user)
	comment := data.CreateTestComment()
	comment.PostID = post.ID
	comment.UserID = user.ID
	rqparm := requestParam{uri: commentURI, auth: token, body: comment}
	code, body := rqparm.postJSON()
	assert.Equal(t, http.StatusOK, code)
	response := &data.Comment{}

	if err := json.Unmarshal(body, response); err != nil {
		t.Error("解析comment json 失败")
	}

	if response.Body != comment.Body {
		t.Error("comment data not true")
	}

	comment.PostID = ""
	rqparm = requestParam{uri: commentURI, auth: token, body: comment}
	code, _ = rqparm.postJSON()
	assert.NotEqual(t, http.StatusOK, code)

	comment.UserID = ""
	rqparm = requestParam{uri: commentURI, auth: token, body: comment}
	code, _ = rqparm.postJSON()
	assert.NotEqual(t, http.StatusOK, code)
}

func TestUpdateComment(t *testing.T) {
	_, _, comment, token := createTestComment()
	comment.Body = "hello world"
	rqparm := requestParam{auth: token, body: comment, uri: commentURI}
	code, body := rqparm.putJSON()
	assert.Equal(t, http.StatusOK, code)
	response := &data.Comment{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Error("update comment data error")
	}
	if response.Body != comment.Body {
		t.Error("update comment data not true")
	}
	comment.Body = "he"
	rqparm = requestParam{auth: token, body: comment, uri: commentURI}
	code, _ = rqparm.putJSON()
	assert.NotEqual(t, http.StatusOK, code)
}

func TestDeleteCommentAPI(t *testing.T) {
	_, _, comment, token := createTestComment()
	rqparm := requestParam{auth: token, body: comment, uri: commentURI + comment.ID + "1"}
	code, _ := rqparm.deleteJSON()
	assert.NotEqual(t, http.StatusOK, code)

	rqparm = requestParam{auth: token, body: comment, uri: commentURI + comment.ID}
	code, _ = rqparm.deleteJSON()
	assert.Equal(t, http.StatusOK, code)

	rqparm = requestParam{auth: token, body: comment, uri: commentURI + comment.ID}
	code, _ = rqparm.deleteJSON()
	assert.NotEqual(t, http.StatusOK, code)
}

func TestGetCommentBYID(t *testing.T) {
	_, _, comment, token := createTestComment()
	rqparm := requestParam{auth: token, body: comment, uri: commentURI + comment.ID}
	code, body := rqparm.getJSON()
	assert.Equal(t, http.StatusOK, code)

	response := &data.Comment{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Error("get comment error")
	}
	if response.ID != comment.ID {
		t.Error("get comment not match")
	}
}

func createTestComment() (data.User, data.Post, data.Comment, string) {
	user, token := createTestUserAndToken()
	post := data.CreateTestPost(user)
	post.GeneratePost(user)
	comment := data.CreateTestComment()
	comment.CreateComment(user, post)

	return user, post, comment, token
}
