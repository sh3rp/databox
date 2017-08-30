package db

import (
	"crypto/sha256"
	"strings"

	"github.com/sh3rp/databox/msg"
)

type InMemorySearchEngine struct {
	TermIndex        map[string][]*msg.Link
	LinkTagSignature map[string][]byte
	LinkTagList      map[string][]string
}

func NewInMemorySearchEngine() SearchEngine {
	return &InMemorySearchEngine{
		TermIndex:        make(map[string][]*msg.Link),
		LinkTagSignature: make(map[string][]byte),
		LinkTagList:      make(map[string][]string),
	}
}

func (se *InMemorySearchEngine) ReIndex() {

}

func (se *InMemorySearchEngine) IndexLink(link *msg.Link) error {

	tags := normalizeTags(link)

	if _, ok := se.LinkTagSignature[link.Id]; ok {
		changed := len(tags) != len(se.LinkTagSignature[link.Id]) &&
			!compareHashes(hashTags(tags), se.LinkTagSignature[link.Id])

		if changed {
			se.UnIndexLink(link)
		}
	}

	for _, t := range tags {
		se.TermIndex[t] = append(se.TermIndex[t], link)
	}

	se.LinkTagSignature[link.Id] = hashTags(tags)
	se.LinkTagList[link.Id] = tags

	return nil
}
func (se *InMemorySearchEngine) UnIndexLink(link *msg.Link) error {
	for _, t := range se.LinkTagList[link.Id] {
		if _, ok := se.TermIndex[t]; ok {
			left := 0
			for ; ; left++ {
				if se.TermIndex[t][left].Id == link.Id {
					break
				}
			}
			leftArray := se.TermIndex[t][:left]
			rightArray := se.TermIndex[t][left+1:]
			se.TermIndex[t] = leftArray
			se.TermIndex[t] = append(se.TermIndex[t], rightArray...)
		}
	}
	return nil
}
func (se *InMemorySearchEngine) FindLinks(term string) []*msg.Link {
	return se.TermIndex[term]
}

func normalizeTags(link *msg.Link) []string {
	var tags []string
	for _, t := range link.Tags {
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

func compareHashes(a, b []byte) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func hashTags(tags []string) []byte {
	hasher := sha256.New()
	for _, t := range tags {
		hasher.Write([]byte(t))
	}
	return hasher.Sum(nil)
}
