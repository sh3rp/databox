package search

import (
	"crypto/sha256"
	"strings"

	"github.com/emirpasic/gods/sets/treeset"
	"github.com/sh3rp/databox/msg"
)

type SearchEngine interface {
	Index(msg.Key, []string) error
	UnIndex(msg.Key) error
	Find(string, int, int) []msg.Key
	FindPartial(string, int, int) []msg.Key
}

type Node struct {
	Char   byte
	Nodes  map[byte]*Node
	Leaves []msg.Key
}

func NormalizeTags(inTags []string) []string {
	if len(inTags) == 0 {
		return nil
	}
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

	var tags []string
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
