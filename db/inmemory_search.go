package db

import (
	"strings"

	"github.com/sh3rp/databox/msg"
)

type InMemorySearchEngine struct {
	TermIndex map[string][]*msg.Link
}

func NewInMemorySearchEngine() SearchEngine {
	return &InMemorySearchEngine{
		TermIndex: make(map[string][]*msg.Link),
	}
}

func (se *InMemorySearchEngine) IndexLink(link *msg.Link) error {

	tags := getTags(link)

	for _, t := range tags {
		se.TermIndex[t] = append(se.TermIndex[t], link)
	}

	return nil
}
func (se *InMemorySearchEngine) UnIndexLink(*msg.Link) error {
	return nil
}
func (se *InMemorySearchEngine) FindLinks(term string) []*msg.Link {
	return se.TermIndex[term]
}

func getTags(link *msg.Link) []string {
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
