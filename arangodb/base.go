package arango

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/yangioc/bk_pack/log"
	"github.com/yangioc/bk_pack/util"
)

var _instans *Manager

var NoDataError = errors.New("data not find")

type Manager struct {
	Client driver.Database
}

func LaunchInstans(addr, username, password, database string) {
	_instans = &Manager{}
	db, err := conn(addr, username, password, database)
	if err != nil {
		panic(err)
	}
	_instans.Client = db
}

func conn(addr, username, password, database string) (driver.Database, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{addr},
	})
	if err != nil {
		log.Errorf("[Arango][New] Http new connection error, err: %v", err)
		return nil, err
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(username, password),
	})
	if err != nil {
		log.Errorf("[Arango][New] Driver new client error, err: %v", err)
		return nil, err
	}

	db, err := c.Database(context.TODO(), database)
	if err != nil {
		log.Errorf("[Arango][New] Client database error, database: %v, err: %v", database, err)
		return nil, err
	}

	log.Infof("[Arango][New] Connect success, address: %v, database: %v", addr, database)
	return db, nil
}

// Quary String
func (self *Manager) Quary(ctx context.Context, query string, bindVars map[string]interface{}, outDoc interface{}) error {
	cursor, err := self.Client.Query(ctx, query, bindVars)
	if err != nil {
		return err
	}
	defer cursor.Close()

	if !cursor.HasMore() {
		return nil
	}

	docs := []interface{}{}
	for {
		doc := make(map[string]interface{})
		_, err = cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		}
		docs = append(docs, doc)
	}

	// 避免driver撈出科學記號數字變成浮點數，過一層轉換將金額欄位維持整數格式
	bytes, err := util.Marshal(docs)
	if err != nil {
		return err
	}
	err = util.Unmarshal(bytes, outDoc)
	if err != nil {
		return err
	}

	return err
}

// Quary String
func (self *Manager) QuaryMap(ctx context.Context, query string, bindVars map[string]interface{}, outDoc map[string]interface{}) error {
	cursor, err := self.Client.Query(ctx, query, bindVars)
	if err != nil {
		return err
	}
	defer cursor.Close()

	if !cursor.HasMore() {
		return nil
	}

	for {
		doc := make(map[string]interface{})
		meta, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		}
		outDoc[meta.Key] = doc
	}

	return err
}

// Quary Create
func (self *Manager) Create(ctx context.Context, collection string, doc interface{}) error {
	coll, err := self.Client.Collection(ctx, collection)
	if err != nil {
		return err
	}

	_, err = coll.CreateDocument(ctx, doc)
	return err
}

func (self *Manager) CreateAndResKey(ctx context.Context, collection string, doc interface{}) (string, error) {
	coll, err := self.Client.Collection(ctx, collection)
	if err != nil {
		return "", err
	}

	meta, err := coll.CreateDocument(ctx, doc)
	return meta.Key, err
}

// Quary Read
func (self *Manager) Read(ctx context.Context, collection string, key string, outDoc interface{}) error {
	coll, err := self.Client.Collection(ctx, collection)
	if err != nil {
		return err
	}
	_, err = coll.ReadDocument(ctx, key, outDoc)
	return err
}

// Quary Reads
func (self *Manager) Reads(ctx context.Context, collection string, keys []string, outDoc interface{}) error {
	coll, err := self.Client.Collection(ctx, collection)
	if err != nil {
		return err
	}
	_, sliceErr, err := coll.ReadDocuments(ctx, keys, outDoc)
	fmt.Println(sliceErr, err)
	return err
}

// Quary Update
func (self *Manager) Update(ctx context.Context, collection string, key string, doc interface{}) error {
	coll, err := self.Client.Collection(ctx, collection)
	if err != nil {
		return err
	}
	_, err = coll.UpdateDocument(ctx, key, doc)
	return err
}

// Quary Delete
func (self *Manager) Delete(ctx context.Context, collection string, key string) error {
	coll, err := self.Client.Collection(ctx, collection)
	if err != nil {
		return err
	}
	_, err = coll.RemoveDocument(ctx, key)
	return err
}

// Transaction
func (self *Manager) Transaction(ctx context.Context, action string, opt *driver.TransactionOptions, outDoc interface{}) error {
	res, err := self.Client.Transaction(ctx, action, opt)
	switch reflect.TypeOf(outDoc).Kind() {
	case reflect.Ptr, reflect.Struct:
		jsByte, _ := json.Marshal(res)
		err := json.Unmarshal(jsByte, outDoc)
		if err != nil {
			return err
		}

	}
	return err
}
