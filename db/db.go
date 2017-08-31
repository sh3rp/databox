package db

import "github.com/sh3rp/databox/msg"

type BoxDB interface {
	NewBox(string, string) (*msg.Box, error)
	SaveBox(*msg.Box) error
	GetBoxById(string) (*msg.Box, error)
	GetBoxes() ([]*msg.Box, error)
	DeleteBox(string) error

	NewLink(name string, url string, boxId string) (*msg.Link, error)
	SaveLink(*msg.Link) error
	GetLinkById(string) (*msg.Link, error)
	GetLinks() ([]*msg.Link, error)
	GetLinksByBoxId(string) ([]*msg.Link, error)
	DeleteLink(string) error

	//	NewNote(name string, text []byte, boxId string) (*msg.Note, error)
	//	SaveNote(*msg.Note) error
	//	GetNoteById(string) (*msg.Note, error)
	//	GetNotes() ([]*msg.Note, error)
	//	GetNotesByBoxId(string) ([]*msg.Note, error)
	//	DeleteNote(string) error
}
