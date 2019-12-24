package msg

//卡劵
type HandlerCard interface {
	CardPassCheck()			//卡券审核通过
	CardNotPassCheck()		//卡券审核失败
	UserGetCard()			//用户领取卡券
	UserDelCard()			//用户删除卡券
	UserConsumeCard()		//用户核销卡券
	UserViewCard()			//用户浏览会员卡
	MerchantOrder()			//微小店用户下单付款
}