package etcd

import (
	"time"
	"context"
	"github.com/coreos/etcd/clientv3"
)


// Options options
type Options struct {
	Endpoints    []string `json:"endpoints,omitempty" yaml:"endpoints"`
	Username     string `json:"username,omitempty" yaml:"username"`
	Password     string `json:"password,omitempty" yaml:"password"`
	
	TimeOut			 time.Duration
	// Client    	 *clientv3.Client
	// Kv       		 clientv3.KV
}


// Option Options function
type Option func(*Options)


// NewOptions new Options
func NewOptions(opts ...Option) Options {
	opt := Options{
		Endpoints:    []string{"127.0.0.1"},
		Username:     "root",
		Password:     "root",
		TimeOut:  	2 * time.Second , // s
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}


// 定义key变更事件常量
const (
	KeyCreateChangeEvent = iota
	KeyUpdateChangeEvent
	KeyDeleteChangeEvent
)

// key 变化事件
type KeyChangeEvent struct {
	Type  int
	Key   string
	Value []byte
}

// 监听key 变化响应
type WatchKeyChangeResponse struct {
	Event      chan *KeyChangeEvent
	CancelFunc context.CancelFunc
	Watcher    clientv3.Watcher
}

type TxResponse struct {
	Success bool
	LeaseID clientv3.LeaseID
	Lease   clientv3.Lease
	Key     string
	Value   string
}