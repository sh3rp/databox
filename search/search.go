package search

import (
	"crypto/sha256"
	"strings"

	"github.com/emirpasic/gods/sets/treeset"
	"github.com/sh3rp/databox/db"
)

type SearchEngine interface {
	Index(db.Key, []string) error
	UnIndex(db.Key) error
	Find(string, int, int) []db.Key
	FindPartial(string, int, int) []db.Key
}

type Node struct {
	Char   byte
	Nodes  map[byte]*Node
	Leaves []db.Key
}

func NormalizeTags(inTags []string) []string {
	var tags []string
	set := treeset.NewWithStringComparator()
	for _, t := range inTags {
		if strings.Contains(t, " ") {
			tempTags := strings.Split(t, " ")
			for _, tt := range tempTags {
				set.Add(tt)
			}
		} else {
			set.Add(t)
		}
	}
	for _, v := range set.Values() {
		tags = append(tags, v.(string))
	}
	return tags
}

func CompareHashes(a, b []byte) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func HashTags(tags []string) []byte {
	hasher := sha256.New()
	for _, t := range tags {
		hasher.Write([]byte(t))
	}
	return hasher.Sum(nil)
}
