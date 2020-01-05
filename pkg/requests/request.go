package requests

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
	"time"
	"strconv"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string, timeout time.Duration) (string, string, error, bool) {

    // 超时时间：5秒
    client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := client.Get(url)
	var isTimeout bool
    if err != nil {
        isTimeout = true
	}else{
		isTimeout = false
	}
	status_code:= resp.StatusCode //获取返回状态码
    defer resp.Body.Close()
    body, err2 := ioutil.ReadAll(resp.Body)

	return string(body),strconv.Itoa(status_code),err2,isTimeout                                                                                               
    
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json   application/x-www-form-unlencoded
// content：     请求放回的内容
func Post(url string, data interface{}, timeout time.Duration) (string, string, error, bool) {
    contentType := "application/x-www-form-unlencoded"
    // 超时时间：5秒
    client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
    jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	var isTimeout bool
	if err != nil {
        isTimeout = true
	}else{
		isTimeout = false
	}
    defer resp.Body.Close()

    result, err2 := ioutil.ReadAll(resp.Body)
    return string(result),err2.Error(),err2,isTimeout
}