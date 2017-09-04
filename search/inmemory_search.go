package search

import (
	"errors"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/sh3rp/databox/msg"
)

type InMemorySearchEngine struct {
	TermIndex        map[string]*hashset.Set
	LinkTagSignature map[msg.Key][]byte
	LinkTagList      map[msg.Key][]string
	searchTree       *Node
}

func NewInMemorySearchEngine() SearchEngine {
	return &InMemorySearchEngine{
		TermIndex:        make(map[string]*hashset.Set),
		LinkTagSignature: make(map[msg.Key][]byte),
		LinkTagList:      make(map[msg.Key][]string),
		searchTree: &Node{
			Nodes: make(map[byte]*Node),
		},
	}
}

func (se *InMemorySearchEngine) Index(id msg.Key, inTags []string) error {
	if len(inTags) == 0 {
		return nil
	}

	if id.Id == "" || id.BoxId == "" {
		return errors.New("No such id")
	}

	tags := NormalizeTags(inTags)

	if _, ok := se.LinkTagSignature[id]; ok {
		changed := len(tags) != len(se.LinkTagSignature[id]) &&
			!CompareHashes(HashTags(tags), se.LinkTagSignature[id])

		if changed {
			se.UnIndex(id)
		}
	}

	for _, t := range tags {
		if se.TermIndex[t] == nil {
			se.TermIndex[t] = hashset.New()
		}
		se.TermIndex[t].Add(id)
	}

	se.LinkTagSignature[id] = HashTags(tags)
	se.LinkTagList[id] = tags

	return nil
}

func (se *InMemorySearchEngine) UnIndex(id msg.Key) error {
	for _, t := range se.LinkTagList[id] {
		if _, ok := se.TermIndex[t]; ok {
			se.TermIndex[t].Remove(id)
		}
	}

	delete(se.LinkTagSignature, id)
	delete(se.LinkTagList, id)

	return nil
}

func (se *InMemorySearchEngine) Find(term string, limit, page int) []msg.Key {
	termList := se.TermIndex[term]

	var keys []msg.Key
	for _, t := range termList.Values() {
		keys = append(keys, t.(msg.Key))
	}

	if len(keys) <= limit {
		return keys
	} else {
		return keys[0 : limit+1]
	}
}

func (se *InMemorySearchEngine) FindPartial(term string, limit, page int) []msg.Key {
	return nil
}
