package common

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"io/ioutil"
	"net/http"
	"wx-api/config"
	"wx-api/tools"
)

//微信入口绑定
//signature	微信加密签名，signature结合了开发者填写的token参数和请求中的timestamp参数、nonce参数。
//timestamp	时间戳
//nonce	随机数
//校验请求是否来源微信
func WxEventListen(w http.ResponseWriter,r *http.Request) {
	values  := r.URL.Query()
	echostr	:= values.Get("echostr")

	//存在echostr时为接入验证
	if echostr != "" {
		signature := values.Get("signature")
		timestamp := values.Get("timestamp")
		nonce     := values.Get("nonce")
		token := config.WxC.Token
		//将token，timestamp，nonce放入slice，进行字典排序并sha1加密
		rescode := tools.HashStrings([]string{
			timestamp,
			nonce,
			token,
		})
		fmt.Println(signature,rescode)
		//校验加密字符是否与signature相同
		if signature == rescode {
			fmt.Fprint(w,echostr)
			return
		}
	} else {
		//接收xml数据包
		res,_ := ioutil.ReadAll(r.Body)
		//data := string(res)
		//fmt.Println(data)
		//解析xml
		resData := make(tools.XmlMap)
		err := xml.Unmarshal(res,&resData)
		Handler(resData)
		fmt.Println(err)
	}
}


//微信获取AccessToken 返回指定微信公众号的at信息
func GetAccessToken () (token string,err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	var key = "WX_TOKEN"
	//将 token 存入 redis （微信token有7200s 有效期）
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s",config.RedisC.Host,config.RedisC.Port),
		Password: config.RedisC.PassWord, // no password set
		DB:       0,  // use default DB
	})

	token, err = client.Get("key").Result()
	if  err != nil || token == ""  {
		url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential"
		url = fmt.Sprintf( "%s&appid=%s&secret=%s",url,config.WxC.AppId,config.WxC.AppSecret)
		fmt.Println(url)
		resp,err := WxHttpRequest("get",url,nil)
		if err != nil {
			panic("Token Get Failed:"+err.Error())
		}

		tokenInfo := make(map[string]interface{})
		err = json.Unmarshal([]byte(resp),&tokenInfo)
		if err != nil {
			panic("Json To Map Failed:"+err.Error())
		}

		token,ok := tokenInfo["access_token"]
		if ok {
			err = client.Set(key, token, 7100).Err()
			if err != nil {
				panic("Redis Set Failed"+err.Error())
			}
		} else {
			panic(tokenInfo)
		}
	}
	return token,err
}


//获取授权页ticket
func GetWxTicket(){
	wxTicket("WX_JS_TICKET")
}

//获取jsapi ticket
func GetWxJsApiTicket(){
	wxTicket("WX_TICKET")
}

//获取微信ticket
func wxTicket(key string) string {
	var getType string
	switch key {
		case "WX_TICKET":
			getType = "wx_card"
			break
		case "WX_JS_TICKET":
			getType = "jsapi"
			break

	}
	//将 token 存入 redis （微信token有7200s 有效期）
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s",config.RedisC.Host,config.RedisC.Port),
		Password: config.RedisC.PassWord, // no password set
		DB:       0,  // use default DB
	})

	token, err := client.Get("key").Result()
	if err != nil || token == "" {
		//获取token
		ACCESS_TOKEN,_ := GetAccessToken()
		url := "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
		url = fmt.Sprintf( "?access_token=%s&%s&appid=%s&secret=%s&type=%s",url,ACCESS_TOKEN,config.WxC.AppId,config.WxC.AppSecret,getType)
		fmt.Println(url)
		token,err := WxHttpRequest("get",url,nil)
		if err != nil {
			panic("Token Get Failed:"+err.Error())
		}
		err = client.Set(key, token, 7100).Err()
		if err != nil {
			panic(err)
		}
	}
	return token
}

//微信提交API方法
func WxHttpRequest(method,url string,param map[string]string) (data string,err error){
	switch method {
		case "get":
			data,err = tools.Get(url)
			break
		case "post":
			data,err = tools.Post(url,param)
			break
		default:
			err = errors.New("Request type does not exist")
	}
	return data,err
}

