package ldap

type Options struct {
	Server 		  string   `yaml:"server"` //HOST:PORT
	BaseDN 		  string   `yaml:"baseDN"`
	BindDn   	  string   `yaml:"bindDn"`
	BindPass  	  string   `yaml:"bindPass"`
	AuthFilter    string   `yaml:"authFilter"`
	AllAuthFilter string   `yaml:"allAuthFilter"`
	SearchField   []string `yaml:"searchField"`
}

// Option Options function
type Option func(*Options)

// NewOptions new Options
func NewOptions(opts ...Option) Options {
	opt := Options{
		Server:     "127.0.0.1:389",
		BaseDN:      "test",
		BindDn:      "test",
		BindPass:    "test",
		AuthFilter:  "test",
		AllAuthFilter: "test",
		SearchField: []string{"test"},
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}


 func Server(server string) Option {
	return func(o *Options) {
		o.Server = server
	}
}

 func BaseDN(baseDN string) Option {
	return func(o *Options) {
		o.BaseDN = baseDN
	}
}

 func BindDn(bindDn string) Option {
	return func(o *Options) {
		o.BindDn = bindDn
	}
}

 func BindPass(bindPass string) Option {
	return func(o *Options) {
		o.BindPass = bindPass
	}
}

 func AuthFilter(authFilter string) Option {
	return func(o *Options) {
		o.AuthFilter = authFilter
	}
}

 func AllAuthFilter(allAuthFilter string) Option {
	return func(o *Options) {
		o.AllAuthFilter = allAuthFilter
	}
}

 func SearchField(searchField []string) Option {
	return func(o *Options) {
		o.SearchField = searchField
	}
}
