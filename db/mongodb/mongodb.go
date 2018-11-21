package mongodb

import (
	"github.com/wangbokun/go/log"
	"gopkg.in/mgo.v2"
)

type MongoDBConfig struct {
	Url string `yaml:"url"`
	DB  string `yaml:"db"`
}

var collection *mgo.Database

func MongoInit(url, db string) error {
	mgo.SetDebug(true)
	sess, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	collection = sess.DB(db)
	return nil
}

func Upsert(collect string, index map[string]interface{}, doc interface{}) (*mgo.ChangeInfo, error) {
	info, err := collection.C(collect).Upsert(index, doc)
	return info, err
}

func Query(collect string, selector interface{}, results interface{}) error {
	err := collection.C(collect).Find(selector).Sort("-_id").All(results)
	return err
}

func Delete(collect string, selector interface{}) (*mgo.ChangeInfo, error) {
	info, err := collection.C(collect).RemoveAll(selector)
	log.Debug("Mongo delete[%s], selector=%#v; info=%#v, err=%v", collect, selector, info, err)
	return info, err
}

func DeleteId(collect string, selector interface{}) error {
	err := collection.C(collect).RemoveId(selector)
	log.Debug("Mongo deleteId[%s], selector=%#v; err=%v", collect, selector, err)
	return err
}
