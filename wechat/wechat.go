package wechat


import ( 
	"github.com/wangbokun/go/http/client" 
	"github.com/wangbokun/go/types"
	"github.com/wangbokun/go/log"
	
)

type WechatToken struct {
	AccessToken string `json:"access_token"`
}

type weChatMessage struct {
	Text    weChatMessageContent `yaml:"text,omitempty" json:"text,omitempty"`
	ToUser  string               `yaml:"touser,omitempty" json:"touser,omitempty"`
	ToParty string               `yaml:"toparty,omitempty" json:"toparty,omitempty"`
	Totag   string               `yaml:"totag,omitempty" json:"totag,omitempty"`
	AgentID string               `yaml:"agentid,omitempty" json:"agentid,omitempty"`
	Safe    string               `yaml:"safe,omitempty" json:"safe,omitempty"`
	Type    string               `yaml:"msgtype,omitempty" json:"msgtype,omitempty"`
}

type weChatMessageContent struct {
	Content string `json:"content"`
}

type weChatResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}


func SendMsg(user, msgType, ctx,token,agentId string){
	
	msg := &weChatMessage{
		Text: weChatMessageContent{
			Content: ctx,
		},
		ToUser:  user,
		ToParty: "",
		Totag:   "",
		AgentID: agentId,
		Type:    msgType, // msgType: text taskcard  markdown ...
		Safe:    "0",
	} 
 
	s,err:=types.ToString(msg)
	if err!=nil{
		log.Error("to string faild, Error: %s",err )
	}

 	resp, error := client.Post(wxSendTextUrl + token,"application/json;charset=utf-8",s)
	if error!=nil{
		log.Error("wechat send  faild, Error: %s",error )
	}

	log.Debug("wechat send  status: %s",resp)
}