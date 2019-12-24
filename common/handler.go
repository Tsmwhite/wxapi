package common

import (
	"fmt"
	"time"
)


var eventTypeAll = []string{
	"subscribe",
	"unsubscribe",
	"SCAN",
	"LOCATION",
	"CLICK",
	"VIEW",
}

var msgTypeAll = []string{
	"text",
	"image",
	"voice",
	"video",
	"shortvideo",
	"location",
	"link",
}

//当前时间
var NOW_TIME  = time.Now().Unix()

func Handler(data map[string]string){
	if msgType,ok := data["MsgType"]; ok {
		if msgType == "Event"{
			//事件推送


		}else{
			//消息推送

		}
	} else {
		fmt.Println("消息类型未注册",data)
	}
}

