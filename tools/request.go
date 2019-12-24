package tools

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

//http post 请求 content-Type
const (
	C_URL_ENCODE = "application/x-www-form-urlencoded"
	C_JSON	     = "application/json"
	C_TEXT	     = "text/xml"
	C_FILE  	 = "multipart/form-data"
)

//默认 post 请求 content-Type
var PostContentType = C_JSON

//发起http get 请求
func Get(url string) (data string,err error){
	res,err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body,err := ioutil.ReadAll(res.Body)
	data = string(body)
	if data == "" {
		err = errors.New("Data does not exist")
	}
	fmt.Println(data)
	return
}

//发起http post 请求
func Post(url string,param map[string]string) (data string,err error) {
	//遍历参数生成请求string 例："test=abc"
	req := ""
	for v,k := range param {
		req += fmt.Sprintf("%s=%s",k,v)
	}
	if req == "" {
		err = errors.New("Post request parameter cannot be empty")
		return
	}
	res, err := http.Post(url,PostContentType,strings.NewReader(req))
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	data = string(body)
	if data == "" {
		err = errors.New("Data does not exist")
	}
	fmt.Println(data)
	return
}

//sha1加密
func HashStrings(sl []string) string {
	sort.Strings(sl)
	h := sha1.New()
	for _, s := range sl {
		io.WriteString(h, s)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}