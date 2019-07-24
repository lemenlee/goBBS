package apitest

import (
	"bbs/api"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
)

func requestJSON(method string, uri string, param interface{}) (int, []byte) {
	router := api.SetupRouter()
	// 将参数转化为json比特流
	jsonByte, _ := json.Marshal(param)

	// 构造请求，json数据以请求body的形式传递
	req := httptest.NewRequest(method, uri, bytes.NewReader(jsonByte))

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

func postJSON(uri string, param interface{}) (int, []byte) {
	return requestJSON("POST", uri, param)
}

func putJSON(uri string, param interface{}) (int, []byte) {
	return requestJSON("PUT", uri, param)
}

func requestHeader(method string, uri string, param string) (int, []byte) {
	router := api.SetupRouter()
	req := httptest.NewRequest(method, uri, nil)
	req.Header.Set("Authorization", param)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	// 提取响应
	result := w.Result()
	defer result.Body.Close()
	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return w.Code, body
}

func putHeader(uri string, param string) (int, []byte) {
	return requestHeader("PUT", uri, param)
}
