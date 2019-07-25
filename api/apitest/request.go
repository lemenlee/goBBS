package apitest

import (
	"bbs/api"
	"bbs/data"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
)

type requestParam struct {
	auth   string
	body   interface{}
	uri    string
	method string
}

func (rqParam *requestParam) requestJSON() (int, []byte) {
	router := api.SetupRouter()
	// 将参数转化为json比特流
	jsonByte, _ := json.Marshal(rqParam.body)

	// 构造请求，json数据以请求body的形式传递
	req := httptest.NewRequest(rqParam.method, rqParam.uri, bytes.NewReader(jsonByte))
	if len(rqParam.auth) > 0 {
		req.Header.Set("Authorization", rqParam.auth)
	}
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// assert.Equal(t, 200, )
	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return w.Code, body
}

func (rqParam *requestParam) postJSON() (int, []byte) {
	rqParam.method = "POST"
	return rqParam.requestJSON()
}

func (rqParam *requestParam) putJSON() (int, []byte) {
	rqParam.method = "PUT"
	return rqParam.requestJSON()
}

func (rqParam *requestParam) getJSON() (int, []byte) {
	rqParam.method = "GET"
	return rqParam.requestJSON()
}

func (rqParam *requestParam) deleteJSON() (int, []byte) {
	rqParam.method = "DELETE"
	return rqParam.requestJSON()
}

func createTestUser() data.User {
	data.Db.Delete(&data.User{})
	user := data.CreateUser("lemon@qq.com", "123456", "lemonlee")
	user.GeneratePasswordHash()
	user.CreateUserDB()
	return user
}

func createTestUserAndToken() (data.User, string) {
	data.Db.Delete(&data.User{})
	user := data.CreateUser("lemon@qq.com", "123456", "lemonlee")
	user.GeneratePasswordHash()
	user.CreateUserDB()

	token := data.JwtToken{}
	token.GenerateToken(user.UserLogin)
	return user, token.AccessTokenStr
}
