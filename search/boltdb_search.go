package search

import "github.com/boltdb/bolt"

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
	})
	return &BoltSearchEngine{DB: boltDB}
}

func (s *BoltSearchEngine) Index(id Key, inTags []string) error {

	tags := NormalizeTags(inTags)

	if _, ok := se.LinkTagSignature[id]; ok {
		changed := len(tags) != len(se.LinkTagSignature[id]) &&
			!CompareHashes(HashTags(tags), se.LinkTagSignature[id])

		if changed {
			se.UnIndex(id)
		}
	}

	for _, t := range tags {
		se.TermIndex[t] = append(se.TermIndex[t], id)
	}

	se.LinkTagSignature[id] = HashTags(tags)
	se.LinkTagList[id] = tags

	return nil
}

func (s *BoltSearchEngine) UnIndex(key Key) error {
	return nil
}

func (s *BoltSearchEngine) Find(term string, count, page int) []Key {
	return nil
}

func (s *BoltSearchEngine) FindPartial(term string, count, page int) []Key {
	return nil
}

func (s *BoltSearchEngine) insertTermSig(id Key, sig []byte) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(TAG_SIG_BUCKET))
		return bucket.Put([]byte(id), sig)
	})
}

func (s *BoltSearchEngine) getTermSig(id Key) ([]byte, error) {
	var sig []byte

}
