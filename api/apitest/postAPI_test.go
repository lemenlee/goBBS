package apitest

import (
	"bbs/data"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const postURI = "/api/v1/post/"

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	// shutdown()
	os.Exit(code)
}

func setup() {
	fmt.Println("----------test begin------------")
	data.Db.Delete(&data.Post{})
}

// func shutdown() {
// 	fmt.Println("----------test end------------")
// }

func TestCreatePost(t *testing.T) {
	_, tokenStr := createTestUserAndToken()
	post := data.PostModel{Body: "123412313123132131"}
	rqparm := requestParam{uri: postURI, auth: tokenStr, body: post}
	code, body := rqparm.postJSON()
	assert.Equal(t, http.StatusOK, code)
	response := &data.Post{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}
	if len(response.ID) == 0 || response.Body != post.Body {
		t.Errorf("post create error")
	}

	post.Body = ""
	rqparm = requestParam{uri: postURI, auth: tokenStr, body: post}
	code, body = rqparm.postJSON()
	assert.Equal(t, http.StatusForbidden, code)

}

func TestGetPost(t *testing.T) {
	user, tokenStr := createTestUserAndToken()
	post := data.CreateTestPost(user)
	post.GeneratePost(user)
	rqparm := requestParam{uri: postURI + post.ID, auth: tokenStr}
	code, body := rqparm.getJSON()
	assert.Equal(t, http.StatusOK, code)
	response := &data.Post{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}
	if response.ID != post.ID || response.Body != post.Body {
		t.Errorf("get post data error")
	}

	rqparm.uri += "1"
	code, body = rqparm.getJSON()
	assert.NotEqual(t, http.StatusOK, code)
}

func TestGetAllPost(t *testing.T) {
	user, tokenStr := createTestUserAndToken()
	for i := 0; i < 10; i++ {
		post := data.CreateTestPost(user)
		post.GeneratePost(user)
	}
	rqparm := requestParam{uri: postURI, auth: tokenStr}
	code, body := rqparm.getJSON()
	assert.Equal(t, http.StatusOK, code)
	response := &[]data.Post{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}
	// fmt.Println(*response)
	if len(*response) != 10 {
		t.Errorf("posts count not true")
	}
}

func TestDeletePost(t *testing.T) {
	user, tokenStr := createTestUserAndToken()
	post := data.CreateTestPost(user)
	post.GeneratePost(user)

	rqparm := requestParam{uri: postURI + post.ID, auth: tokenStr}
	code, _ := rqparm.deleteJSON()
	assert.Equal(t, http.StatusOK, code)

	post1, _ := data.FindPostByID(post.ID)
	fmt.Println(post1)
	if post1.ID == post.ID {
		t.Errorf("post delete fail")
	}

	code, _ = rqparm.deleteJSON()
	assert.Equal(t, http.StatusForbidden, code)
}

func TestUpdatePost(t *testing.T) {
	user, tokenStr := createTestUserAndToken()
	post := data.CreateTestPost(user)
	post.GeneratePost(user)

	post.Body = "hello world"
	rqparm := requestParam{body: post, auth: tokenStr, uri: postURI}
	code, _ := rqparm.putJSON()
	assert.Equal(t, http.StatusOK, code)

	post1, err := data.FindPostByID(post.ID)
	if err != nil {
		t.Error("after update post, find post error")
	}
	if post1.Body != "hello world" {
		t.Error("update post body error")
	}

	post.Body = "hel"
	rqparm.body = post
	code, _ = rqparm.putJSON()
	assert.NotEqual(t, http.StatusOK, code)

}
