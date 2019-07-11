package ldap

import (
	"fmt"
	"log"
	"strings"
	"go-common/project/ops/OMS/libs/types"

	"gopkg.in/ldap.v2"
	// "go-common/project/ops/prometheusWebHook/models/httpclient"
)

func SnsLdapSearch(Username, Server string, Port int, Bind_dn, Bind_password, Search_base_dns, AuthFilter string, Attributes []string, ctx string) {

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", Server, Port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind(Bind_dn, Bind_password)
	if err != nil {
		log.Fatal(err)
	}

	searchRequest := ldap.NewSearchRequest(
		Search_base_dns,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(AuthFilter, Username),
		Attributes,
		nil,
	)

	sr, err := l.Search(searchRequest)
	// fmt.Println(sr.Entries)

	for _, v := range sr.Entries {
		for _, vv := range v.Attributes {

			if vv.Name == "mobile" {
				log.Println("Alarm sending User ", Username, " Phone No is", strings.SplitN(vv.Values[0], " ", -1)[2], "content is", ctx)
		
			}
		}
	}

	if err != nil {
		log.Fatal(err)
	}
	if len(sr.Entries) != 1 {
		fmt.Println("user is not exist")
	}
}

type data struct{ 
	Attributes  Attributes
}

type Attributes struct{
	mobile  string `json:"mobile"`
	mail  	string `json:"mail"`
	cn  	string `json:"cn"`
}

func LdapSearch(Username, Server string, Port int, Bind_dn, Bind_password, Search_base_dns, AuthFilter string, Attributes []string) (name,Phone string){

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", Server, Port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind(Bind_dn, Bind_password)
	if err != nil {
		log.Fatal(err)
	}

	searchRequest := ldap.NewSearchRequest(
		Search_base_dns,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(AuthFilter, Username),
		Attributes,
		nil,
	)

	sr, err := l.Search(searchRequest)

	for _, v := range sr.Entries {
		attributes := Attributes
		error := types.Format(v.Attributes, &attributes)
		if error != nil{
			fmt.Println(error)
		}
		var info string
		for _, vv := range v.Attributes {
			if vv.Name == "mobile" {
				info = info + strings.SplitN(vv.Values[0], " ", -1)[2] 
				log.Println("Alarm sending User ", Username, " Phone No is", strings.SplitN(vv.Values[0], " ", -1)[2])
				return Username,strings.SplitN(vv.Values[0], " ", -1)[2]
			}else{
				info = info+vv.Values[0]
			}

		}
		fmt.Println("info################>:",info)	
	}
		
	if err != nil {
		log.Fatal(err)
	}
	if len(sr.Entries) != 1 {
		fmt.Println("user is not exist")
	}
	return
}