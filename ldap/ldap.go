package ldap

import (
	"fmt"
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
