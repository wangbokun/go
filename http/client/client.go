package client

import ( 
	"io/ioutil" 
	"net/http"
	"strings"
)

func Post(Url, ContentType , Content string) (body []byte,err error){
	// ContentType :  application/x-www-form-urlencoded  or  application/json;charset=utf-8
	resp, err := http.Post(Url,
		ContentType,
		strings.NewReader(Content))
	if err != nil { 
		return nil, err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil { 
		return nil, err
	}

 	return body, nil
}


func Get(Url string) (body []byte,err error){

	// if strings.HasPrefix(Url, "http://") == false{
	// 	Url = "http://"+Url
	// }
	resp, err := http.Get(Url)

	if err != nil {
		return nil,err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	return body,nil
}