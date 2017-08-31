package db

import (
	"bytes"
	"encoding/gob"

	"github.com/boltdb/bolt"
	"github.com/sh3rp/databox/msg"
)

var BOX_BUCKET = "box"
var LINK_BUCKET = "link"

type BoltDB struct {
	DB *bolt.DB
}

func NewBoltDB(name, dir string) BoxDB {
	boltDB, err := bolt.Open(dir+"/"+name, 0644, nil)

	return &BoltDB{DB: boltDB}
}

func (db *BoltDB) NewBox(name string, description string) (*msg.Box, error) {
	box := &msg.Box{
		Id:          GenerateID(),
		Name:        name,
		Description: description,
	}
	err := db.insertObject(BOX_BUCKET, box.Id, box)
	return box, err
}

func (db *BoltDB) SaveBox(box *msg.Box) error {
	return nil
}

func (db *BoltDB) GetBoxById(id string) (*msg.Box, error) {
	return nil, nil
}

func (db *BoltDB) GetBoxes() ([]*msg.Box, error) {
	return nil, nil
}

func (db *BoltDB) DeleteBox(id string) error {
	return nil
}

func (db *BoltDB) NewLink(name string, url string, boxId string) (*msg.Link, error) {
	return nil, nil
}

func (db *BoltDB) SaveLink(link *msg.Link) error {
	return nil
}

func (db *BoltDB) GetLinkById(id string) (*msg.Link, error) {
	return nil, nil
}

func (db *BoltDB) GetLinks() ([]*msg.Link, error) {
	return nil, nil
}

func (db *BoltDB) GetLinksByBoxId(id string) ([]*msg.Link, error) {
	return nil, nil
}

func (db *BoltDB) DeleteLink(id string) error {
	return nil
}

// bolt specific elements

func (db *BoltDB) insertObject(bucket string, id string, obj interface{}) error {
	var data bytes.Buffer
	err := gob.NewEncoder(&data).Encode(obj)

	if err != nil {
		return err
	}

	db.insertKV([]byte(bucket), []byte(id), data.Bytes())

	return nil
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
	err := db.DB.View(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(table)

		if err != nil {
			return err
		}

		bytes := bucket.Get(key)
		data = &kv{K: key, V: bytes}
		return nil
	})
	return data, err
}

func (db *BoltDB) getKVs(table, key []byte) ([]*kv, error) {
	var data []*kv
	err := db.DB.View(func(tx *bolt.Tx) error {
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
