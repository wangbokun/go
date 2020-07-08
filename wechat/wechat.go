package wechat


import ( 
	"github.com/wangbokun/go/http/client" 
	"github.com/wangbokun/go/types"
	"github.com/wangbokun/go/log"
	
)

type WechatToken struct {
	AccessToken string `json:"access_token"`
}

type textWeChatMessage struct {
	Text        textMessageContent 				 `yaml:"text,omitempty" json:"text,omitempty"`
	// Taskcard  	taskcardMessageContent  		 `yaml:"text,omitempty" json:"text,omitempty"`
	ToUser      string               			 `yaml:"touser,omitempty" json:"touser,omitempty"`
	ToParty 	string               			 `yaml:"toparty,omitempty" json:"toparty,omitempty"`
	Totag   	string              			 `yaml:"totag,omitempty" json:"totag,omitempty"`
	AgentID 	string               		   	 `yaml:"agentid,omitempty" json:"agentid,omitempty"`
	Safe    	string              			 `yaml:"safe,omitempty" json:"safe,omitempty"`
	Type   	 	string              		 	 `yaml:"msgtype,omitempty" json:"msgtype,omitempty"`
}

type markdownWeChatMessage struct {
	Markdown    markdownMessageContent 				 `yaml:"markdown,omitempty" json:"markdown,omitempty"`
	ToUser      string               			 `yaml:"touser,omitempty" json:"touser,omitempty"`
	ToParty 	string               			 `yaml:"toparty,omitempty" json:"toparty,omitempty"`
	Totag   	string              			 `yaml:"totag,omitempty" json:"totag,omitempty"`
	AgentID 	string               		   	 `yaml:"agentid,omitempty" json:"agentid,omitempty"`
	Safe    	string              			 `yaml:"safe,omitempty" json:"safe,omitempty"`
	Type   	 	string              		 	 `yaml:"msgtype,omitempty" json:"msgtype,omitempty"`
}

type markdownMessageContent struct {
	Content string `json:"content"`
}

type textMessageContent struct {
	Content string `json:"content"`
}
 

type weChatResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}


func SendMsg(user, mstype, ctx, token, agentId string) error {
	switch mstype {
	case "markdown":
		msg := &markdownWeChatMessage{
			Markdown: markdownMessageContent{
				Content: ctx,
			},
			ToUser:  user,
			ToParty: "",
			Totag:   "",
			AgentID: agentId,
			Type:    "markdown", // msgType: text taskcard  markdown ...
			Safe:    "0",
		} 
		s,err:=types.ToString(msg)
		if err!=nil{
			log.Error("to string faild, Error: %s",err )
			return err
		}
		if err = send(s,token); err !=nil{
			return err
		}
	default:
		msg := &textWeChatMessage{
			Text: textMessageContent{
				Content: ctx,
			},
			ToUser:  user,
			ToParty: "",
			Totag:   "",
			AgentID: agentId,
			Type:    "text", // msgType: text taskcard  markdown ...
			Safe:    "0",
		} 
		s,err:=types.ToString(msg)
		if err!=nil{
			log.Error("to string faild, Error: %s",err )
			return err
		}
		if err = send(s,token); err !=nil{
			return err
		}
	}
	return nil
}

func send(s,token string) error {
	_, err := client.Post(wxSendTextUrl + token,"application/json;charset=utf-8",s)
	if err !=nil{
		log.Error("wechat send  faild, Error: %s", err )
		return err
	}
	// log.Debug("wechat send  status: %s",resp)
	return nil
}