package wechat

import ( 
	"github.com/wangbokun/go/http/client"
	"github.com/wangbokun/go/types"
)

const (
	wechartUrl = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid="
)
type Token struct{
	Errcode 		int 		`json:"errcode"`
	Errmsg  		string 		`json:"errmsg"`
	Access_token 	string 		`json:"access_token"`
	Expires_in 		int 		`json:"expires_in"`
}

func GetToken(corpid,corpsecret string) (string,error) {
	Url:=wechartUrl+corpid+"&corpsecret="+corpsecret
	ctx,err:=client.Get(Url)
	if err != nil{
		return "err",err
	}
	WechatToken := &Token{}

	err =types.FormatJSON(ctx,WechatToken)
	if err != nil{
		return "err",err
	}
	return WechatToken.Access_token,nil
}