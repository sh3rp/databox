package db

import "github.com/sh3rp/databox/msg"

type SearchEngine interface {
	IndexLink(*msg.Link) error
	UnIndexLink(*msg.Link) error
	FindLinks(string) []*msg.Link
	ReIndex()
}
