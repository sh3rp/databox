package db

import (
	"bytes"
	"encoding/gob"
	"errors"

	"github.com/boltdb/bolt"
	"github.com/sh3rp/databox/msg"
)

var BOX_BUCKET = "box"
var LINK_BUCKET = "link"

type BoltDB struct {
	DB *bolt.DB
}

func NewBoltDB(dbfilename string) BoxDB {
	boltDB, err := bolt.Open(dbfilename, 0600, nil)
	if err != nil {
		panic(err) // kill the program, we need a DB to run
	}
	return &BoltDB{DB: boltDB}
}

func (db *BoltDB) NewBox(name string, description string) (*msg.Box, error) {
	if name == "" {
		return nil, errors.New("Must supply a name")
	}

	if description == "" {
		return nil, errors.New("Must supply a description")
	}

	box := &msg.Box{
		Id:          GenerateID(),
		Name:        name,
		Description: description,
	}
	return box, db.SaveBox(box)
}

func (db *BoltDB) SaveBox(box *msg.Box) error {
	return db.insertBox(box)
}

func (db *BoltDB) GetBoxById(id string) (*msg.Box, error) {
	obj, err := db.getBox(BOX_BUCKET, id)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (db *BoltDB) GetBoxes() ([]*msg.Box, error) {
	return db.getAllBoxes()
}

func (db *BoltDB) DeleteBox(id string) error {
	return db.deleteKey([]byte(BOX_BUCKET), []byte(id))
}

func (db *BoltDB) NewLink(name string, url string, boxId string) (*msg.Link, error) {
	if boxId == "" {
		return nil, errors.New("Must supply a box ID")
	}

	_, err := db.GetBoxById(boxId)

	if err != nil {
		return nil, err
	}

	link := &msg.Link{
		Id:    GenerateID(),
		Name:  name,
		Url:   url,
		BoxId: boxId,
	}

	return link, db.SaveLink(link)
}

func (db *BoltDB) SaveLink(link *msg.Link) error {
	return db.insertLink(link)
}

func (db *BoltDB) GetLinkById(box, id string) (*msg.Link, error) {
	obj, err := db.getLink(box, id)

	return obj, err
}

func (db *BoltDB) GetLinksByBoxId(id string) ([]*msg.Link, error) {
	_, err := db.GetBoxById(id)

	if err != nil {
		return nil, err
	}

	return db.getAllLinks(id)
}

func (db *BoltDB) DeleteLink(boxId, id string) error {
	_, err := db.GetBoxById(boxId)

	if err != nil {
		return err
	}

	return db.deleteKey([]byte(boxId), []byte(id))
}

// bolt specific elements

func (db *BoltDB) keyFromLink(link *msg.Link) string {
	return "l-" + link.Id
}

func (db *BoltDB) insertBox(box *msg.Box) error {
	var data bytes.Buffer
	err := gob.NewEncoder(&data).Encode(box)

	if err != nil {
		return err
	}

	err = db.insertKV([]byte(BOX_BUCKET), []byte(box.Id), data.Bytes())

	return err
}

func (db *BoltDB) insertLink(link *msg.Link) error {
	var data bytes.Buffer
	err := gob.NewEncoder(&data).Encode(link)

	if err != nil {
		return err
	}

	db.insertKV([]byte(link.BoxId), []byte("l-"+link.Id), data.Bytes())

	return nil
}

func (db *BoltDB) getBox(bucket string, id string) (*msg.Box, error) {
	var buf bytes.Buffer

	kv, err := db.getKV([]byte(bucket), []byte(id))

	if err != nil {
		return nil, err
	}

	if len(kv.V) == 0 {
		return nil, errors.New("key does not exist")
	}

	_, err = buf.Write(kv.V)

	if err != nil {
		return nil, err
	}
	obj := &msg.Box{}
	err = gob.NewDecoder(&buf).Decode(obj)

	if err != nil {
		return nil, err
	}

	return obj, err
}

func (db *BoltDB) getLink(box string, id string) (*msg.Link, error) {
	var buf bytes.Buffer

	kv, err := db.getKV([]byte(box), []byte("l-"+id))

	if err != nil {
		return nil, err
	}

	if len(kv.V) == 0 {
		return nil, errors.New("key does not exist")
	}

	_, err = buf.Write(kv.V)

	if err != nil {
		return nil, err
	}

	obj := &msg.Link{}
	err = gob.NewDecoder(&buf).Decode(obj)

	if err != nil {
		return nil, err
	}

	return obj, err
}

func (db *BoltDB) getAllBoxes() ([]*msg.Box, error) {
	var objs []*msg.Box

	allKvs, err := db.getAllKVs([]byte(BOX_BUCKET))

	if err != nil {
		return nil, err
	}

	for _, kv := range allKvs {
		buf := &bytes.Buffer{}
		_, err = buf.Write(kv.V)
		obj := &msg.Box{}
		err = gob.NewDecoder(buf).Decode(obj)
		objs = append(objs, obj)
	}

	return objs, nil
}

func (db *BoltDB) getAllLinks(box string) ([]*msg.Link, error) {
	var objs []*msg.Link

	allKvs, err := db.getAllKVs([]byte(box))

	if err != nil {
		return nil, err
	}

	for _, kv := range allKvs {
		buf := &bytes.Buffer{}
		_, err = buf.Read(kv.V)
		obj := &msg.Link{}
		err = gob.NewDecoder(buf).Decode(obj)
		objs = append(objs, obj)
	}

	return objs, nil
}

func (db *BoltDB) getObjects(bucket, prefix string) ([]interface{}, error) {
	var objs []interface{}

	allKvs, err := db.getKVs([]byte(bucket), []byte(prefix))

	if err != nil {
		return nil, err
	}

	for _, kv := range allKvs {
		var obj interface{}
		buf := &bytes.Buffer{}
		_, err = buf.Read(kv.V)
		err = gob.NewDecoder(buf).Decode(obj)
		objs = append(objs, obj)
	}

	return objs, nil
}

type kv struct {
	K []byte
	V []byte
}

func (db *BoltDB) insertKV(table, key, value []byte) error {
	err := db.DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(table)

		if err != nil {
			return err
		}

		return bucket.Put(key, value)
	})
	return err
}

func (db *BoltDB) getKV(table, key []byte) (*kv, error) {
	var data *kv
	err := db.DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(table)

		if err != nil {
			return err
		}

		bytes := bucket.Get(key)

		if len(bytes) == 0 {
			return errors.New("Key does not exist")
		}

		data = &kv{K: key, V: bytes}

		return nil
	})
	return data, err
}

func (db *BoltDB) getKVs(table, key []byte) ([]*kv, error) {
	var data []*kv
	err := db.DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(table)

		if err != nil {
			return err
		}

		c := bucket.Cursor()

		for k, v := c.Seek(key); k != nil && bytes.HasPrefix(k, key); k, v = c.Next() {
			data = append(data, &kv{K: k, V: v})
		}

		return nil
	})
	return data, err
}

func (db *BoltDB) getAllKVs(table []byte) ([]*kv, error) {
	var data []*kv
	err := db.DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(table)

		if err != nil {
			return err
		}

		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			data = append(data, &kv{K: k, V: v})
		}

		return nil
	})

	return data, err
}

func (db *BoltDB) deleteKey(table, key []byte) error {
	err := db.DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(table)

		if err != nil {
			return err
		}

		return bucket.Delete(key)
	})
	return err
}
