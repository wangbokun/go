package ldap

import (
	"fmt"
<<<<<<< HEAD
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
=======
	"gopkg.in/ldap.v2"
)

type LDAP struct {
	Server 		 string `yaml:"server"` //HOST:PORT
	BaseDN 		 string `yaml:"baseDN"`
	BindDn   	 string `yaml:"bindDn"`
	BindPass  	 string `yaml:"bindPass"`
	AuthFilter   string `yaml:"authFilter"`	
}

func (l *LDAP) Search(search string) (*ldap.SearchResult, error) {
	conn, err := ldap.Dial("tcp", l.Server)
	if err != nil {
		return nil, err
	}
	searchRequest := ldap.NewSearchRequest(
		l.BaseDN, // The base dn to search
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		search,               //"(cn=*)" The filter to apply
		[]string{"dn", "cn"}, // A list attributes to retrieve
		nil,
	)
	res, err := conn.Search(searchRequest)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (l *LDAP) Auth(username, password string) error {
	conn, err := ldap.Dial("tcp", l.Server)
	if err != nil {
		return err
	}
	defer conn.Close()
 
	err = conn.Bind(l.BindDn, l.BindPass)
	if err != nil {
		return err
	}

	searchRequest := ldap.NewSearchRequest(
		l.BaseDN, //binddn, //"dc=example,dc=com", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		// fmt.Sprintf("(&(objectClass=organizationalPerson)(sAMAccountName=%s))", username),
		fmt.Sprintf(l.AuthFilter, username),
		[]string{"dn","cn"}, // A list attributes to retrieve
		nil,
	) 
	
	sr, err := conn.Search(searchRequest)
	if err != nil {
		return err
	} 

	fmt.Println(sr)
	if len(sr.Entries) != 1 {
		return fmt.Errorf("user is not exist")
	}
	fmt.Println(sr.Entries[0].DN,password)
	if err := conn.Bind(sr.Entries[0].DN, password); err != nil {
		return err
	}
	return nil
}
>>>>>>> master
