package msg


//卡劵
type HandlerMsg interface {
	MsgText()			//文本格式
	MsgImage()			//图片格式
	MsgVoice()			//声音
	MsgVideo()			//视频
	MsgShortVideo()		//小视频
	MsgLocation()		//上传地理位置
	MsgLink()			//链接
}

