package korm

import (
	"fmt"
    "database/sql"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"github.com/wangbokun/go/log"
	"github.com/wangbokun/go/codec"
)


// // MySQL mysql
type MySQL struct {
	opts Options
	db *sql.DB
    Eloquent *gorm.DB
}


// New file config
func New(opts ...Option) *MySQL {
	options := NewOptions(opts...)
	return &MySQL{
		opts: options,
	}
}



// Init init
func (my *MySQL) Init(opts ...Option) {
	for _, o := range opts {
		o(&my.opts)
	}
}


func (my *MySQL)Connect() error {

	// "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8", my.opts.Username, my.opts.Password, my.opts.Hostname, my.opts.Port, my.opts.Database)
    Eloquent, err := my.Open(my.opts.DbType,dsn)
	Eloquent.LogMode(true)

	if err != nil {
		log.Error("mysql connect error %v", err)
        return err
	} else {
		log.Info("mysql connect success!")
	}

	if Eloquent.Error != nil {
		log.Error("database error %v", Eloquent.Error)
	}
    return nil
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
