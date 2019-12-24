package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"wx-api/common"
	"wx-api/config"
	"wx-api/tools"
)

func main() {
	wxXml := `<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1348831860</CreateTime>
  <MsgType><![CDATA[text]]></MsgType>
  <Content><![CDATA[this is a test]]></Content>
  <MsgId>1234567890123456</MsgId>
</xml>`

	xmlM := make(tools.XmlMap)
	xml.Unmarshal([]byte(wxXml),&xmlM)
	fmt.Println(xmlM)


	wxc := config.WxConfig{
		AppId:     "*****",
		AppSecret: "*****",
		AesKey:    "*****",
		Token:     "*****",
	}
	rc := config.RedisConfig{
		Host:     "127.0.0.1",
		Port:     "6379",
		PassWord: "",
	}
	router := gin.Default()
	router.GET("/wxEvent", func(c *gin.Context) {
		fmt.Println(c.Params)
		common.WxEventListen(c.Writer,c.Request)
	})

	config.WxInitConfig(wxc,rc)
	token,err := common.GetAccessToken()
	fmt.Println(token,err)
	err = router.Run(":80")
	fmt.Println(err)

}