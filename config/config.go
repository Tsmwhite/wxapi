package config

//微信公众号配置
type WxConfig struct {
	AppId 		string
	AppSecret	string
	AesKey 		string
	Token 		string
}
//redis 配置
type RedisConfig struct {
	Host 		string
	Port 		string
	PassWord	string
}

var WxC WxConfig
var RedisC RedisConfig

func WxInitConfig (wxc WxConfig,rc RedisConfig) {
	WxC = wxc
	RedisC = rc
}



