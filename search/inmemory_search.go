package search

import (
	"crypto/sha256"
	"strings"
)

type InMemorySearchEngine struct {
	TermIndex        map[string][]Key
	LinkTagSignature map[Key][]byte
	LinkTagList      map[Key][]string
	searchTree       *Node
}

func NewInMemorySearchEngine() SearchEngine {
	return &InMemorySearchEngine{
		TermIndex:        make(map[string][]Key),
		LinkTagSignature: make(map[Key][]byte),
		LinkTagList:      make(map[Key][]string),
		searchTree: &Node{
			Nodes: make(map[byte]*Node),
		},
	}
}

func (se *InMemorySearchEngine) ReIndex() {

}

func (se *InMemorySearchEngine) Index(id Key, inTags []string) error {

	tags := normalizeTags(inTags)

	if _, ok := se.LinkTagSignature[id]; ok {
		changed := len(tags) != len(se.LinkTagSignature[id]) &&
			!compareHashes(hashTags(tags), se.LinkTagSignature[id])

		if changed {
			se.UnIndex(id)
		}
	}

	for _, t := range tags {
		se.TermIndex[t] = append(se.TermIndex[t], id)

		bytes := []byte(t)
		node := se.searchTree

		for _, b := range bytes {
			if node.Nodes[b] == nil {
				node.Nodes[b] = &Node{Char: b, Nodes: make(map[byte]*Node)}
			}
			node = node.Nodes[b]
		}
		node.Leaves = append(node.Leaves, Key(t))
	}

	se.LinkTagSignature[id] = hashTags(tags)
	se.LinkTagList[id] = tags

	return nil
}

func (se *InMemorySearchEngine) UnIndex(id Key) error {
	for _, t := range se.LinkTagList[id] {
		if _, ok := se.TermIndex[t]; ok {
			left := 0
			for ; ; left++ {
				if se.TermIndex[t][left] == id {
					break
				}
			}
			leftArray := se.TermIndex[t][:left]
			rightArray := se.TermIndex[t][left+1:]
			se.TermIndex[t] = leftArray
			se.TermIndex[t] = append(se.TermIndex[t], rightArray...)
		}
	}

	delete(se.LinkTagSignature, id)
	delete(se.LinkTagList, id)

	return nil
}

func (se *InMemorySearchEngine) Find(term string, limit int) []Key {
	termList := se.TermIndex[term]

	if len(termList) <= limit {
		return termList
	} else {
		return termList[0 : limit+1]
	}
}

func (se *InMemorySearchEngine) FindPartial(term string, limit int) []Key {
	return nil
}

func normalizeTags(inTags []string) []string {
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
