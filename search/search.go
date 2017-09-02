package search

import (
	"crypto/sha256"
	"strings"
)

type SearchEngine interface {
	Index(Key, []string) error
	UnIndex(Key) error
	Find(string, int, int) []Key
	FindPartial(string, int, int) []Key
}

type Key string

type Node struct {
	Char   byte
	Nodes  map[byte]*Node
	Leaves []Key
}

func NormalizeTags(inTags []string) []string {
	var tags []string
	for _, t := range inTags {
		if strings.Contains(t, " ") {
			tempTags := strings.Split(t, " ")
			for _, tt := range tempTags {
				tags = append(tags, tt)
			}
		} else {
			tags = append(tags, t)
		}
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
