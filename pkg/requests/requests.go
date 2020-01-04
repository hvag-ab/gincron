package requests

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"io/ioutil"
	."fmt"//省略写fmt
	"strings"
	"net/url" 
	"encoding/json"
	"bytes"
	
)



func HttpGet(urls string) string{
	resp, err := http.Get(urls)
	if err != nil {
		Println(err)
	}
	body, err2 := ioutil.ReadAll(resp.Body)
	if err != nil {
		Println(err2)
	}
	defer resp.Body.Close()

	return string(body)
}

func HttpPost(urls,data string) string{
	resp, err := http.Post(urls,
		"application/x-www-form-urlencoded",
		strings.NewReader(data))
	if err != nil {
		Println(err)
	}
 
	defer resp.Body.Close()
	body, err2:= ioutil.ReadAll(resp.Body)
	if err2 != nil {
		Println(err)
	}
 
	return string(body)
}
//Tips：使用这个方法的话，第二个参数要设置成”application/x-www-form-urlencoded”，否则post参数无法传递。

//一种是使用http.PostForm方法
func HttpPostForm(urls string) string{
	resp, err := http.PostForm(urls,
		url.Values{"key": {"Value"}, "id": {"123"}})
 
	if err != nil {
		Println(err)
	}
 
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err != nil {
		Println(err2)
	}
 
	return string(body)
 
}


func HttpDo(urls,data string) string{
	client := &http.Client{}
	
	var datas = "hvag=abc" //datas是post的传递的参数
	req, _ := http.NewRequest("POST",urls, strings.NewReader(datas))

 
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")
 
	resp, _ := client.Do(req)
 
	defer resp.Body.Close()
 
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Println(err)
	}
 
	return string(body)
}
//同上面的post请求，必须要设定Content-Type为application/x-www-form-urlencoded，post参数才可正常传递。

func Httpjson(){
	urls := "http://restapi3.apiary.io/notes"
  

    var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
    req, err := http.NewRequest("POST", urls, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    Println("response Status:", resp.Status)
    Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
	Println("response Body:", string(body))
	

	// values := map[string]string{"username": username, "password": password}

	// jsonValue, _ := json.Marshal(values)

	// resp, err := http.Post(authAuthenticatorUrl, "application/json", bytes.NewBuffer(jsonValue))

}

func Httpget(){
	req, err := http.NewRequest("GET", "http://api.themoviedb.org/3/tv/popular", nil)
    if err != nil {
        Print(err)
       
    }

    q := req.URL.Query()
    q.Add("api_key", "key_from_environment_or_flag")
    q.Add("another_thing", "foo & bar")
    req.URL.RawQuery = q.Encode()

    Println(req.URL.String())
    // Output:
    // http://api.themoviedb.org/3/tv/popular?another_thing=foo+%26+bar&api_key=key_from_environment_or_flag
    var resp *http.Response
    resp, err = http.DefaultClient.Do(req)
    if err != nil {
        Print(err)
    }
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	Println("response Body:", string(body))

}



func Json2map(){
	str:="{\"address\":\"北京\",\"username\":\"kongyixueyuan\"}"
	myMap:=make(map[string]string)
	json.Unmarshal([]byte(str),&myMap)
	Println(myMap)

}

func Map2json(){
	user := make(map[string]string)
	user["username"]="kongyixueyuan"
	user["address"]="北京"
	jsonStr, err := json.Marshal(user)
	if err!=nil{
	   Println(err)
	}
	Printf("%s",jsonStr)
}

type Person struct {
	UserName string `json:"username"`
	Age int `json:"age"`
 }
func Struct2json(){
	str:=Person{"hvag",30}
	jsonStr,_:= json.Marshal(str)

	Println(string(jsonStr))
}

type Vih struct{
	name string
	age int
}

func ParseJSON(url string) []Vih{

	var v []Vih
	var s Vih

	body := `{
               "hits":{
				"hits":[
					{
						"hvag":{"name":"hvag","age":23}
					},
					{
						"hvag":{"name":"gavh","age":18}
					}
				]
            }`

	_byte := []byte(body)
	//

	json_data := jsoniter.Get(_byte, "hits","hits")
	//获取第一个
	size := json_data.Size()//获取数组长度
	Println("size=",size)
	_data := []byte(json_data.ToString())
	for i := 0 ; i< size ; i++{
		source := jsoniter.Get(_data,i,"hvag").ToString()
		json.Unmarshal([]byte(source), &s)
		v = append(v,s)
	}
	return v
}