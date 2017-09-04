package search

import (
	"bytes"
	"encoding/gob"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/golang/protobuf/proto"
	"github.com/sh3rp/databox/msg"
)

var TERM_BUCKET = "terms"
var TAG_SIG_BUCKET = "tag_signature"
var LINK_BUCKET = "links"

type BoltSearchEngine struct {
	DB *bolt.DB
}

func NewBoltSearchEngine(dbfilename string) SearchEngine {
	boltDB, err := bolt.Open(dbfilename, 0600, nil)
	if err != nil {
		panic(err) // kill the program, we need a DB to run
	}
	err = boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(TERM_BUCKET))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte(TAG_SIG_BUCKET))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte(LINK_BUCKET))
		if err != nil {
			return err
		}
		return nil
	})
	return &BoltSearchEngine{DB: boltDB}
}

func (s *BoltSearchEngine) Index(id msg.Key, inTags []string) error {
	if len(inTags) == 0 {
		return nil
	}

	tags := NormalizeTags(inTags)

	termSig := s.getTermSig(id)

	changed := len(termSig) == 0 || len(tags) != len(s.getTags(id)) || !CompareHashes(HashTags(tags), termSig)

	if changed {
		s.UnIndex(id)
		for _, t := range tags {
			s.addTermMatch(t, id)
		}

		s.saveTermSig(id, HashTags(tags))
		s.saveTags(id, tags)
	}

	return nil
}

func (s *BoltSearchEngine) UnIndex(key msg.Key) error {
	terms := s.getTags(key)
	for _, t := range terms {
		s.removeTermMatch(t, key)
	}

	return nil
}

func (s *BoltSearchEngine) Find(term string, count, page int) []msg.Key {
	links := s.getTermMatches(term)

	num := len(links)

	if num <= count || num <= (count*page) {
		return links
	} else {
		return links[count*page : count+(count*page)]
	}
}

func (s *BoltSearchEngine) FindPartial(term string, count, page int) []msg.Key {
	return nil
}

func (s *BoltSearchEngine) addTermMatch(term string, key msg.Key) {
	matches := s.getTermMatches(term)
	matches = append(matches, key)
	s.saveTermMatches(term, matches)
}

func (s *BoltSearchEngine) removeTermMatch(term string, key msg.Key) {
	matches := s.getTermMatches(term)
	if len(matches) > 0 {
		var i int
		for ; i < len(matches); i++ {
			if matches[i] == key {
				break
			}
		}
		newMatches := matches[:i]
		newMatches = append(newMatches, matches[i+1:]...)
		s.saveTermMatches(term, newMatches)
	}
}

func (s *BoltSearchEngine) getTermMatches(term string) []msg.Key {
	var linksBytes []byte
	var links []msg.Key
	var buf bytes.Buffer
	s.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(TERM_BUCKET))
		linksBytes = bucket.Get([]byte(term))
		return nil
	})
	buf.Write(linksBytes)
	gob.NewDecoder(&buf).Decode(&links)
	return links
}

func (s *BoltSearchEngine) saveTermMatches(term string, links []msg.Key) {
	s.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(TERM_BUCKET))
		var buf bytes.Buffer
		gob.NewEncoder(&buf).Encode(links)
		bucket.Put([]byte(term), buf.Bytes())
		return nil
	})
}

func (s *BoltSearchEngine) saveTags(id msg.Key, tags []string) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(LINK_BUCKET))
		key, err := proto.Marshal(&id)
		if err != nil {
			return err
		}
		err = bucket.Put(key, []byte(strings.Join(tags, ",")))
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *BoltSearchEngine) getTags(id msg.Key) []string {
	var tagsStr []byte
	s.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(LINK_BUCKET))
		key, err := proto.Marshal(&id)
		if err != nil {
			return err
		}
		tagsStr = bucket.Get(key)
		return nil
	})

	return strings.Split(string(tagsStr), ",")
}

func (s *BoltSearchEngine) saveTermSig(id msg.Key, sig []byte) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(TAG_SIG_BUCKET))
		key, err := proto.Marshal(&id)
		if err != nil {
			return err
		}
		return bucket.Put(key, sig)
	})
}

func (s *BoltSearchEngine) getTermSig(id msg.Key) []byte {
	var sig []byte

	s.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(TAG_SIG_BUCKET))
		key, err := proto.Marshal(&id)
		if err != nil {
			return err
		}
		sig = bucket.Get(key)
		return nil
	})
	return sig
}
