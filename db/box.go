package db

import "github.com/sh3rp/databox/msg"

type BoxDB interface {
	NewBox(name string) (*msg.Box, error)
	SaveBox(*msg.Box) error
	GetBoxById(int64) (*msg.Box, error)
	GetBoxes() ([]*msg.Box, error)
	DeleteBox(int64) error

	NewLink(name string, url string, boxId int64) (*msg.Link, error)
	SaveLink(*msg.Link) error
	GetLinkById(int64) (*msg.Link, error)
	GetLinks() ([]*msg.Link, error)
	GetLinksByBoxId(int64) ([]*msg.Link, error)
	DeleteLink(int64) error
}
