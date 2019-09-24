package wechat

const (
	wxTokenUrl 		= "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid="
	wxSendTextUrl 	= "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token="
	wxGetUserUrl	= "https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=%s"
)