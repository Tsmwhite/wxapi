package common

//微信二维码Api

//二维码ticket参数（GetWxQRCodeTicket）
type GWQRCT_OPTION struct {
	expire_seconds	int		//该二维码有效时间，以秒为单位。 最大不超过2592000（即30天），此字段如果不填，则默认有效期为30秒。
	action_name		string	//二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值
	//action_info		//二维码详细信息
	scene_id		int64	//场景值ID，临时二维码时为32位非0整型，永久二维码时最大值为100000（目前参数只支持1--100000）
	scene_str		string  //场景值ID（字符串形式的ID），字符串类型，长度限制为1到64
}

//默认参数
var DefaultOpion GWQRCT_OPTION = GWQRCT_OPTION{
	expire_seconds: 300,
	action_name:    "QR_SCENE",
	scene_id:       NOW_TIME,
	scene_str:      "",
}

//设置参数
var QrSetOption func() *GWQRCT_OPTION

//获取微信生成二维码ticket
func GetWxQRCodeTicket() {
	//获取token
	//ACCESS_TOKEN,_ := GetAccessToken()
	//请求url
	//url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s",ACCESS_TOKEN)
	//
}