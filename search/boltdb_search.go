package search

import (
	"bytes"
	"encoding/gob"
	"strings"

	"github.com/boltdb/bolt"
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

func (s *BoltSearchEngine) Index(id Key, inTags []string) error {

	tags := NormalizeTags(inTags)

	termSig := s.getTermSig(id)

	if len(termSig) > 0 {
		changed := len(tags) != len(s.getTags(id)) ||
			!CompareHashes(HashTags(tags), termSig)

		if changed {
			s.UnIndex(id)
		}
	}

	for _, t := range tags {
		s.addTermMatch(t, id)
	}

	s.saveTermSig(id, HashTags(tags))
	s.saveTags(id, tags)

	return nil
}

func (s *BoltSearchEngine) UnIndex(key Key) error {
	terms := s.getTags(key)
	for _, t := range terms {
		s.removeTermMatch(t, key)
	}

	return nil
}

func (s *BoltSearchEngine) Find(term string, count, page int) []Key {
	links := s.getTermMatches(term)

	num := len(links)

	if num <= count || num <= (count*page) {
		return links
	} else {
		return links[count*page : count+(count*page)]
	}
}

func (s *BoltSearchEngine) FindPartial(term string, count, page int) []Key {
	return nil
}

func (s *BoltSearchEngine) addTermMatch(term string, key Key) {
	matches := s.getTermMatches(term)
	matches = append(matches, key)
	s.saveTermMatches(term, matches)
}

func (s *BoltSearchEngine) removeTermMatch(term string, key Key) {
	matches := s.getTermMatches(term)
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

func (s *BoltSearchEngine) getTermMatches(term string) []Key {
	var linksBytes []byte
	var links []Key
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

func (s *BoltSearchEngine) saveTermMatches(term string, links []Key) {
	s.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(TERM_BUCKET))
		var buf bytes.Buffer
		gob.NewEncoder(&buf).Encode(links)
		bucket.Put([]byte(term), buf.Bytes())
		return nil
	})
}

func (s *BoltSearchEngine) saveTags(id Key, tags []string) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(LINK_BUCKET))
		err := bucket.Put([]byte(id), []byte(strings.Join(tags, ",")))
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *BoltSearchEngine) getTags(id Key) []string {
	var tagsStr []byte
	s.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(LINK_BUCKET))
		tagsStr = bucket.Get([]byte(id))
		return nil
	})

	return strings.Split(string(tagsStr), ",")
}

func (s *BoltSearchEngine) saveTermSig(id Key, sig []byte) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(TAG_SIG_BUCKET))
		return bucket.Put([]byte(id), sig)
	})
}

func (s *BoltSearchEngine) getTermSig(id Key) []byte {
	var sig []byte

	s.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(TAG_SIG_BUCKET))
		sig = bucket.Get([]byte(id))
		return nil
	})
	return sig
}
