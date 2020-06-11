package etcd

import (
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/wangbokun/go/codec"
	"github.com/wangbokun/go/log"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"context"
)


// etcd
type Etcd struct {
	opts Options
	client *clientv3.Client
	getResponse *clientv3.GetResponse
	txnResponse *clientv3.TxnResponse
}

// New file config
func New(opts ...Option) *Etcd {
	options := NewOptions(opts...)
	return &Etcd{
		opts: options,
	}
}

// Init init
func (e *Etcd) Init(opts ...Option) {
	for _, o := range opts {
		o(&e.opts)
	}
}

// LoadConfig loadconfig
func (e *Etcd) LoadConfig(v interface{}) error {
	return codec.NewJSONCodec().Format(&e.opts, v)
}

// Connect connect
func (e *Etcd) Connect() error {
	var err error
	e.client, err = clientv3.New(clientv3.Config{
		Endpoints:   e.opts.Endpoints,
		DialTimeout: 2 * time.Second,
		Username: e.opts.Username,
    Password: e.opts.Password,
	})
	return err
}

// get value  from a key
func (e *Etcd) Get(key string) (value []byte, err error) {

	ctx, cancelFunc := context.WithTimeout(context.Background(), e.opts.TimeOut)
	defer cancelFunc()

	if e.getResponse, err = e.client.Get(ctx, key); err != nil {
		log.Error("%s",err)
		return
	}

	if len(e.getResponse.Kvs) == 0 {
		return
	}
	value = e.getResponse.Kvs[0].Value
	return
}


// get values from  prefixKey
func (e *Etcd) GetWithPrefixKey(prefixKey string) (keys [][]byte, values [][]byte, err error) {

	ctx, cancelFunc := context.WithTimeout(context.Background(), e.opts.TimeOut)
	defer cancelFunc()

	if e.getResponse, err = e.client.Get(ctx, prefixKey, clientv3.WithPrefix()); err != nil {
			return
	}

	if len(e.getResponse.Kvs) == 0 {
			return
	}

	keys = make([][]byte, 0)
	values = make([][]byte, 0)

	for i := 0; i < len(e.getResponse.Kvs); i++ {
			keys = append(keys, e.getResponse.Kvs[i].Key)
			values = append(values, e.getResponse.Kvs[i].Value)
	}
	return
}

// get values from  prefixKey limit
func (e *Etcd) GetWithPrefixKeyLimit(prefixKey string, limit int64) (keys [][]byte, values [][]byte, err error) {

	ctx, cancelFunc := context.WithTimeout(context.Background(), e.opts.TimeOut)
	defer cancelFunc()

	if e.getResponse, err = e.client.Get(ctx, prefixKey, clientv3.WithPrefix(), clientv3.WithLimit(limit)); err != nil {
			return
	}

	if len(e.getResponse.Kvs) == 0 {
			return
	}

	keys = make([][]byte, 0)
	values = make([][]byte, 0)

	for i := 0; i < len(e.getResponse.Kvs); i++ {
			keys = append(keys, e.getResponse.Kvs[i].Key)
			values = append(values, e.getResponse.Kvs[i].Value)
	}
	return
}

// put a key
func (e *Etcd) Put(key, value string) (err error) {

	ctx, cancelFunc := context.WithTimeout(context.Background(), e.opts.TimeOut)
	defer cancelFunc()

	if _, err = e.client.Put(ctx, key, value); err != nil {
			return
	}
	return
}

// put a key not exist
func (e *Etcd) PutNotExist(key, value string) (success bool, oldValue []byte, err error) {

	var (
			txnResponse *clientv3.TxnResponse
	)
	ctx, cancelFunc := context.WithTimeout(context.Background(), e.opts.TimeOut)
	defer cancelFunc()

	txn := e.client.Txn(ctx)

	txnResponse, err = txn.If(clientv3.Compare(clientv3.Version(key), "=", 0)).
			Then(clientv3.OpPut(key, value)).
			Else(clientv3.OpGet(key)).
			Commit()

	if err != nil {
			return
	}

	if txnResponse.Succeeded {
			success = true
	} else {
			oldValue = make([]byte, 0)
			oldValue = txnResponse.Responses[0].GetResponseRange().Kvs[0].Value
	}
	return
}

//update
func (e *Etcd) Update(key, value, oldValue string) (success bool, err error) {

	ctx, cancelFunc := context.WithTimeout(context.Background(), e.opts.TimeOut)
	defer cancelFunc()

	txn := e.client.Txn(ctx)

	e.txnResponse, err = txn.If(clientv3.Compare(clientv3.Value(key), "=", oldValue)).
			Then(clientv3.OpPut(key, value)).
			Commit()

	if err != nil {
			return
	}

	if e.txnResponse.Succeeded {
			success = true
	}

	return
}

//del key
func (e *Etcd) Delete(key string) (err error) {

	ctx, cancelFunc := context.WithTimeout(context.Background(), e.opts.TimeOut)
	defer cancelFunc()

	_, err = e.client.Delete(ctx, key)

	return
}

// delete the keys  with prefix key
func (e *Etcd) DeleteWithPrefixKey(prefixKey string) (err error) {

	ctx, cancelFunc := context.WithTimeout(context.Background(), e.opts.TimeOut)
	defer cancelFunc()

	_, err = e.client.Delete(ctx, prefixKey, clientv3.WithPrefix())

	return
}


// watch a key
func (e *Etcd) Watch(key string) (keyChangeEventResponse *WatchKeyChangeResponse) {

	watcher := clientv3.NewWatcher(e.client)
	watchChans := watcher.Watch(context.Background(), key)

	keyChangeEventResponse = &WatchKeyChangeResponse{
			Event:   make(chan *KeyChangeEvent, 250),
			Watcher: watcher,
	}

	go func() {
			for ch := range watchChans {
					if ch.Canceled {
							goto End
					}
					for _, event := range ch.Events {
							e.handleKeyChangeEvent(event, keyChangeEventResponse.Event)
					}
			}

	End:
			log.Info("the watcher lose for key: %s", key)
	}()

	return
}
// handle the key change event
func (e *Etcd) handleKeyChangeEvent(event *clientv3.Event, events chan *KeyChangeEvent) {

	changeEvent := &KeyChangeEvent{
			Key: string(event.Kv.Key),
	}
	switch event.Type {

	case mvccpb.PUT:
			if event.IsCreate() {
					changeEvent.Type = KeyCreateChangeEvent
			} else {
					changeEvent.Type = KeyUpdateChangeEvent
			}
			changeEvent.Value = event.Kv.Value
	case mvccpb.DELETE:

			changeEvent.Type = KeyDeleteChangeEvent
	}
	events <- changeEvent
}

// Close close connect
func (e *Etcd) Close() error {
	return e.client.Close()
}
