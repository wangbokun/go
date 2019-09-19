package wechat


import (
	"fmt"
	"net/http"
	"go-common/project/ops/prometheusWebHook/models/httpclient"
	"go-common/project/ops/prometheusWebHook/models/ldap"
	"go-common/project/ops/OMS/libs/types"
	"go-common/project/ops/prometheusWebHook/models/conf"
	"io/ioutil"
	"github.com/luopengift/gohttp"
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


func WechatSendMsg(user,ctx,token,apiUrl,agentId string){
	
	msg := &weChatMessage{
		Text: weChatMessageContent{
			Content: ctx,
		},
		ToUser:  user,
		ToParty: "",
		Totag:   "",
		AgentID: agentId,
		Type:    "text",
		Safe:    "0",
	}

	postMessageURL := apiUrl + "message/send?access_token=" + token
 
	s,err:=types.ToString(msg)
	if err!=nil{
		log.Error("to string faild, Error: %s",err )
	}
 
	resp, error := gohttp.NewClient().Url(postMessageURL).Body(s).Header("Content-Type", "application/json;charset=utf-8").Post()
	
	if error!=nil{
		log.Error("wechat send  post faild, Error: %s",error )
	}

	log.Debug("wechat send  status: %s",resp.String() )
}


type weChatUser struct {
	Userid  	string               `yaml:"userid,omitempty" json:"userid,omitempty"`
	Name 		string               `yaml:"name,omitempty" json:"name,omitempty"`
	Email		string               `yaml:"email,omitempty" json:"email,omitempty"`
	Mobile   	string               `yaml:"mobile,omitempty" json:"mobile,omitempty"`
	Department  string               `yaml:"department,omitempty" json:"department,omitempty"`
}


func WechatWriteUserHandler(w http.ResponseWriter, r *http.Request){
// func WechatWriteUserHandler(user,name,mobile,token,apiUrl string){
	b, err := ioutil.ReadAll(r.Body)
	
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	m, err := types.BytesToMap(b) 
	if err != nil {
		fmt.Println(err)
	}

	config := &conf.ConfigInfo{}
	err = types.ParseConfigFile("./etc/webhook.json", config)
	if err != nil {
		log.Error("%s",err)
	}

	userInfo := conf.WechatUser{}
	error := types.Format(m, &userInfo)

	if error != nil {
		fmt.Println(error)
	} 

	Username := userInfo.User + "@nio.com"
	
	userMail,mobile :=ldap.LdapSearch(
		Username, config.Addr, config.Port, config.BindDn, config.BindPass,
		config.BaseDn, config.AuthFilter, config.Attributes)
		
	fmt.Println(">>>>>>>>>>",userInfo.User,userMail,mobile)
	token,err := GetToken(config.WechatApiCorpId,config.WechatUserSecret)	
	if err != nil {
		log.Error("%s",err)
	}
	
	msg := &weChatUser{
		Userid: userInfo.User,
		Name: userInfo.User,
		Email: userMail,
		Mobile: mobile,
		Department: "1",
	}

	postMessageURL := config.WechatApiUrl + "user/create?access_token=" + token
 
	s,err:=types.ToString(msg)
	if err!=nil{
		fmt.Println(err)
	}
 
	resp, _ := gohttp.NewClient().Url(postMessageURL).Body(s).Header("Content-Type", "application/json;charset=utf-8").Post()
	fmt.Println(resp.String())
}

func WechatReadUser(user,token,apiUrl string){
	
	URL := apiUrl + "user/get?access_token=" + token +"&userid="+user
	httpclient.HttpClientGet(URL)
}