package db

import (
	"bytes"
	"errors"

	"github.com/boltdb/bolt"
	"github.com/emirpasic/gods/sets/treeset"
	"github.com/golang/protobuf/proto"
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
		Id:          NewBoxKey(),
		Name:        name,
		Description: description,
	}
	return box, db.SaveBox(box)
}

func (db *BoltDB) SaveBox(box *msg.Box) error {
	return db.insertBox(box)
}

func (db *BoltDB) GetBoxById(id msg.Key) (*msg.Box, error) {
	obj, err := db.getBox(id)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (db *BoltDB) GetBoxes() ([]*msg.Box, error) {
	return db.getAllBoxes()
}

func (db *BoltDB) DeleteBox(id msg.Key) error {
	keyData, err := proto.Marshal(&id)

	if err != nil {
		return err
	}

	return db.deleteKey([]byte(BOX_BUCKET), keyData)
}

func (db *BoltDB) NewLink(name string, url string, boxId msg.Key) (*msg.Link, error) {
	if boxId.Id == "" {
		return nil, errors.New("Must supply a box ID")
	}

	box, err := db.GetBoxById(boxId)

	if err != nil {
		return nil, err
	}

	link := &msg.Link{
		Id:   NewLinkKey(box.Id),
		Name: name,
		Url:  url,
	}

	return link, db.SaveLink(link)
}

func (db *BoltDB) SaveLink(link *msg.Link) error {
	return db.insertLink(link)
}

func (db *BoltDB) GetLinkById(id msg.Key) (*msg.Link, error) {
	return db.getLink(id)
}

func (db *BoltDB) GetLinksByBoxId(id msg.Key) ([]*msg.Link, error) {
	box, err := db.GetBoxById(id)

	if err != nil {
		return nil, err
	}

	return db.getAllLinks(*box.Id)
}

func (db *BoltDB) DeleteLink(id msg.Key) error {
	box, err := db.GetBoxById(msg.Key{
		Id:   id.BoxId,
		Type: msg.Key_BOX,
	})

	if err != nil {
		return err
	}
	keyData, err := proto.Marshal(&id)

	return db.deleteKey([]byte(box.Id.Id), keyData)
}

//
// bolt specific elements
//

func (db *BoltDB) insertBox(box *msg.Box) error {
	keyData, err := proto.Marshal(box.Id)

	if err != nil {
		return err
	}

	data, err := proto.Marshal(box)

	if err != nil {
		return err
	}

	err = db.insertKV([]byte(BOX_BUCKET), keyData, data)

	return err
}

func (db *BoltDB) insertLink(link *msg.Link) error {
	var dedupedTags []string
	if len(link.Tags) > 0 {
		set := treeset.NewWithStringComparator()
		for _, t := range link.Tags {
			set.Add(t)
		}
		for _, v := range set.Values() {
			dedupedTags = append(dedupedTags, v.(string))
		}
	}
	link.Tags = dedupedTags

	keyData, err := proto.Marshal(link.Id)

	if err != nil {
		return err
	}

	data, err := proto.Marshal(link)

	db.insertKV([]byte(link.Id.BoxId), keyData, data)

	return nil
}

func (db *BoltDB) getBox(id msg.Key) (*msg.Box, error) {
	keyData, err := proto.Marshal(&id)

	kv, err := db.getKV([]byte(BOX_BUCKET), keyData)

	if err != nil {
		return nil, err
	}

	if len(kv.V) == 0 {
		return nil, errors.New("key does not exist")
	}

	obj := &msg.Box{}
	err = proto.Unmarshal(kv.V, obj)

	return obj, err
}

func (db *BoltDB) getLink(id msg.Key) (*msg.Link, error) {
	keyData, err := proto.Marshal(&id)

	kv, err := db.getKV([]byte(id.BoxId), keyData)

	if err != nil {
		return nil, err
	}

	if len(kv.V) == 0 {
		return nil, errors.New("key does not exist")
	}

	obj := &msg.Link{}
	err = proto.Unmarshal(kv.V, obj)

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
		obj := &msg.Box{}
		err = proto.Unmarshal(kv.V, obj)
		objs = append(objs, obj)
	}

	return objs, nil
}

func (db *BoltDB) getAllLinks(boxId msg.Key) ([]*msg.Link, error) {
	var objs []*msg.Link

	allKvs, err := db.getAllKVs([]byte(boxId.Id))

	if err != nil {
		return nil, err
	}

	for _, kv := range allKvs {
		obj := &msg.Link{}
		err = proto.Unmarshal(kv.V, obj)
		objs = append(objs, obj)
	}

	return objs, nil
}

// byte level

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
