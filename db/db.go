package db

import "github.com/sh3rp/databox/msg"

type BoxDB interface {
	NewBox(string, string) (*msg.Box, error)
	SaveBox(*msg.Box) error
	GetBoxById(msg.Key) (*msg.Box, error)
	GetBoxes() ([]*msg.Box, error)
	DeleteBox(msg.Key) error

	NewLink(string, string, msg.Key) (*msg.Link, error)
	SaveLink(*msg.Link) error
	GetLinkById(msg.Key) (*msg.Link, error)
	GetLinksByBoxId(msg.Key) ([]*msg.Link, error)
	DeleteLink(msg.Key) error

	//	NewNote(name string, text []byte, boxId string) (*msg.Note, error)
	//	SaveNote(*msg.Note) error
	//	GetNoteById(string) (*msg.Note, error)
	//	GetNotes() ([]*msg.Note, error)
	//	GetNotesByBoxId(string) ([]*msg.Note, error)
	//	DeleteNote(string) error
}
