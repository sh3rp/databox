package search

type SearchEngine interface {
	Index(Key, []string) error
	UnIndex(Key) error
	Find(string, int) []Key
	FindPartial(string, int) []Key
	ReIndex()
}

type Key string

type Node struct {
	Char   byte
	Nodes  map[byte]*Node
	Leaves []Key
}
