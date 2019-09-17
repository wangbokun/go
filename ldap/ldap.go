package ldap

import (
	"fmt"
	"github.com/wangbokun/go/codec"
 	"gopkg.in/ldap.v2"
 	"github.com/wangbokun/go/log"
)

type Ldap struct {
	Opts Options
	LdapConn *ldap.Conn
}

// New file config
func New(Opts ...Option) *Ldap {
	options := NewOptions(Opts...)
	return &Ldap{
		Opts: options,
 	}
}

// Init init
func (l *Ldap) Init(Opts ...Option) {
	for _, o := range Opts {
		o(&l.Opts)
	}
}


func (l *Ldap) LoadConfig(v interface{}) error {
	log.Debug("%#v",codec.NewJSONCodec().Format(&l.Opts, v))
	return codec.NewJSONCodec().Format(&l.Opts, v)
}

func (l *Ldap) Connect() (err error){
	l.LdapConn, err = ldap.Dial("tcp", l.Opts.Server)
	if err != nil {
 		return  err
	}
	err = l.LdapConn.Bind(l.Opts.BindDn, l.Opts.BindPass)
	if err != nil {
		log.Error("%s",err) 
	}
	return nil
}
//AuthFilter user: (&(objectClass=organizationalPerson)(sAMAccountName=%s))
//AuthFilter all user: "(&(objectClass=user))"
func (l *Ldap) Search() (*ldap.SearchResult, error) {
	
	// conn, err := ldap.Dial("tcp", l.Opts.Server)
	// if err != nil {
 	// 	return nil, err
	// }
	// defer l.conn.Close()

	// err = conn.Bind(l.Opts.BindDn, l.Opts.BindPass)
	// if err != nil {
	// 	log.Error("%s",err) 
	// } 
  	searchRequest := ldap.NewSearchRequest(
		l.Opts.BaseDN, // The base dn to search
		ldap.ScopeWholeSubtree,ldap.NeverDerefAliases,0,0,false,
		fmt.Sprintf(l.Opts.AuthFilter),             //"(cn=*)" The filter to apply
		l.Opts.SearchField, // A list attributes to retrieve
		nil,
	)
 	// res, err := conn.Search(searchRequest)
	res, err := l.LdapConn.SearchWithPaging(searchRequest,10)
 	if err != nil {
 		return nil, err
	}
	return res, nil

}

func (l *Ldap) Auth(username, password string) error {
 	// conn, err := ldap.Dial("tcp", l.Opts.Server)
	// if err != nil {
 	// 	return err
	// }
	// defer conn.Close()
 
	// err = conn.Bind(l.Opts.BindDn, l.Opts.BindPass)
	// if err != nil {
	// 	return err
	// }
 	searchRequest := ldap.NewSearchRequest(
		l.Opts.BaseDN, //binddn, //"dc=example,dc=com", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		// fmt.Sprintf("(&(objectClass=organizationalPerson)(sAMAccountName=%s))", username),
		fmt.Sprintf(l.Opts.AuthFilter, username),
		[]string{"dn","cn"}, // A list attributes to retrieve
		nil,
	) 
	
	sr, err := l.LdapConn.Search(searchRequest)
	if err != nil {
 		return err
	} 

 	if len(sr.Entries) != 1 {
		return fmt.Errorf("user is not exist")
	}
	fmt.Println(sr.Entries[0].DN,password)
	if err := l.LdapConn.Bind(sr.Entries[0].DN, password); err != nil {
 		return err
	}
	return nil
}


// func (l *Ldap) LdapClose() error {
// 	l.Conn.Close()
// 	return nil
// }

func ChineseName(str string) (s string){
	r := []rune(str)
	strSlice := []string{}
	cnstr := ""
	for i := 0; i < len(r); i++ {
			if r[i] <= 40869 && r[i] >= 19968 {
					cnstr = cnstr + string(r[i])
					strSlice = append(strSlice, cnstr)
			}
	}
	if 0 == len(strSlice) {
		return ""
	}
	return cnstr
}