package korm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"github.com/wangbokun/go/log"
	"strconv"
	"github.com/wangbokun/go/codec"
)

var Eloquent *gorm.DB


type Database interface {
	Open(dbType string, conn string) (db *gorm.DB, err error)
}

// // MySQL mysql
type MySQL struct {
	opts Options
	db *sql.DB
}


// New file config
func New(opts ...Option) *MySQL {
	options := NewOptions(opts...)
	return &MySQL{
		opts: options,
		// Log:  log.DefaultStdLog(),
	}
}



// Init init
func (my *MySQL) Init(opts ...Option) {
	for _, o := range opts {
		o(&my.opts)
	}
}


func (my *MySQL)Connect() {

	// "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	Eloquent, err = db.Open(dbType, 
		fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=%s:%s",
			my.opts.Username, my.opts.Password, my.opts.Database,  my.opts.Hostname, strconv.Itoa(my.opts.Port)
		)
	)
	Eloquent.LogMode(true)

	if err != nil {
		log.Fatalln("mysql connect error %v", err)
	} else {
		log.Println("mysql connect success!")
	}

	if Eloquent.Error != nil {
		log.Fatalln("database error %v", Eloquent.Error)
	}
}


// LoadConfig loadconfig
func (my *MySQL) LoadConfig(v interface{}) error {
	return codec.NewJSONCodec().Format(&my.opts, v)
}


func (*MySQL) Open(dbType string, conn string) (db *gorm.DB, err error) {
	eloquent, err := gorm.Open(dbType, conn)
	return eloquent, err
}

// Close close connect
func (my *MySQL) Close() error {
	return my.db.Close()
}