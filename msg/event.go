package msg

const(
	
)

type EventMsg struct {
	BaseMsg
	Event		string
}

//扫描参数二维码
type ScMsq struct {
	EventMsg
	EventKey	string	//事件KEY值，未订阅时：qrscene_为前缀，后面为二维码的参数值，已订阅时：是一个32位无符号整数，即创建二维码时的二维码scene_id
	Ticket		string	//二维码的ticket，可用来换取二维码图片
}

//上传地理位置
type AddressMsg struct {
	EventMsg
	Latitude	string	//地理位置纬度
	Longitude	string	//地理位置经度
	Precision	string	//地理位置精度
}

//点击自定义菜单
type ClickMsg struct {
	EventMsg
	EventKey	string
}

//事件处理
type HandlerEventer interface {
	Subscribe(*EventMsg)		//关注
	UnSubscribe(*EventMsg)	//取消关注
	Scan(*ScMsq)			//扫描二维码
	Location()		//地理位置
	Click(*ClickMsg)			//点击自定义菜单
	View(*ClickMsg)			//点击菜单跳转链接
}

type HandlerEvent struct {
	WaitEvent chan map[string]string
}

var EventManager = &HandlerEvent{
	make(chan map[string]string),
}

func EventManagerStart()  {
	go func() {
		for {
			if data,ok := <- EventManager.WaitEvent; ok {
				switch data["Event"] {
					case "subscribe":

						break

				}
			}
		}
	}()
}

func (h *HandlerEvent) Subscribe(e *EventMsg){

}

func (h *HandlerEvent) UnSubscribe(e *EventMsg) {

}

func (h *HandlerEvent) Scan(e *ScMsq) {

}

func (h *HandlerEvent) ScaLocation(e *ScMsq) {

}

func (h *HandlerEvent) Click(e *ClickMsg) {

}
func (h *HandlerEvent) View(e *ClickMsg) {

}