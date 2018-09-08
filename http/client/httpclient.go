package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Post(Url, Content string) {

	resp, err := http.Post(Url,
		"application/x-www-form-urlencoded",
		strings.NewReader(Content))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}


func Get(Url string) (ctx interface{},err error){

	if strings.HasPrefix(Url, "http://") == false{
		Url = "http://"+Url
	}
	resp, err := http.Get(Url)

	if err != nil {
		return nil,err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	return body,nil
}