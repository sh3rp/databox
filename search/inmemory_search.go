package search

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

func (se *InMemorySearchEngine) Index(id Key, inTags []string) error {
	if len(inTags) == 0 {
		return nil
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
		se.TermIndex[t] = append(se.TermIndex[t], id)
	}

	se.LinkTagSignature[id] = HashTags(tags)
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

func (se *InMemorySearchEngine) Find(term string, limit, page int) []Key {
	termList := se.TermIndex[term]

	if len(termList) <= limit {
		return termList
	} else {
		return termList[0 : limit+1]
	}
}

func (se *InMemorySearchEngine) FindPartial(term string, limit, page int) []Key {
	return nil
}
