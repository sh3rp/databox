package db

import "github.com/sh3rp/databox/msg"

type BoxDB interface {
	NewBox(string, string, bool) (*msg.Box, error)
	SaveBox(*msg.Box) error
	GetBoxById(string) (*msg.Box, error)
	GetBoxes() ([]*msg.Box, error)
	DeleteBox(string) error
	GetDefaultBox() (*msg.Box, error)

	NewLink(name string, url string, boxId string) (*msg.Link, error)
	SaveLink(*msg.Link) error
	GetLinkById(string) (*msg.Link, error)
	GetLinks() ([]*msg.Link, error)
	GetLinksByBoxId(string) ([]*msg.Link, error)
	DeleteLink(string) error
}
