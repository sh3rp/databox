package db

import (
	"errors"
	"sync"

	"github.com/sh3rp/databox/msg"
)

type InMemoryDB struct {
	boxes      map[string]*msg.Box
	boxesLock  *sync.Mutex
	links      map[string]*msg.Link
	linksLock  *sync.Mutex
	defaultBox *msg.Box
}

func NewInMemoryDB() BoxDB {
	db := InMemoryDB{
		boxes:     make(map[string]*msg.Box),
		links:     make(map[string]*msg.Link),
		boxesLock: new(sync.Mutex),
		linksLock: new(sync.Mutex),
	}
	//db.defaultBox, _ = db.NewBox("default", "Default box", true)
	return &db
}

func (db *InMemoryDB) NewBox(name string, description string, isDefault bool) (*msg.Box, error) {
	if name == "" {
		return nil, errors.New("Name must not be empty")
	}
	if description == "" {
		return nil, errors.New("Description must not be empty")
	}
	box := &msg.Box{
		Id:          GenerateID(),
		Name:        name,
		Description: description,
		IsDefault:   isDefault,
	}
	db.boxesLock.Lock()
	defer db.boxesLock.Unlock()
	db.boxes[box.Id] = box
	if box.IsDefault && db.defaultBox != nil {
		db.boxes[db.defaultBox.Id].IsDefault = false
	}
	db.defaultBox = db.boxes[box.Id]
	return box, nil
}

func (db *InMemoryDB) SaveBox(box *msg.Box) error {
	db.boxesLock.Lock()
	defer db.boxesLock.Unlock()
	if box.Id == "" {
		return errors.New("No id set, cannot save box")
	}
	db.boxes[box.Id] = box
	return nil
}

func (db *InMemoryDB) GetBoxById(id string) (*msg.Box, error) {
	db.boxesLock.Lock()
	defer db.boxesLock.Unlock()
	if box, ok := db.boxes[id]; ok {
		return box, nil
	}
	return nil, errors.New("No such box id")
}

func (db *InMemoryDB) GetBoxes() ([]*msg.Box, error) {
	db.boxesLock.Lock()
	defer db.boxesLock.Unlock()
	var boxes []*msg.Box
	for _, v := range db.boxes {
		boxes = append(boxes, v)
	}
	return boxes, nil
}

func (db *InMemoryDB) DeleteBox(id string) error {
	db.boxesLock.Lock()
	defer db.boxesLock.Unlock()
	delete(db.boxes, id)
	return nil
}

func (db *InMemoryDB) GetDefaultBox() (*msg.Box, error) {
	return db.defaultBox, nil
}

func (db *InMemoryDB) NewLink(name string, url string, boxId string) (*msg.Link, error) {
	if boxId == "" {
		return nil, errors.New("Cannot specify empty box ID")
	}

	_, err := db.GetBoxById(boxId)

	if err != nil {
		return nil, err
	}

	link := &msg.Link{
		Id:    GenerateID(),
		Name:  name,
		Url:   url,
		BoxId: boxId,
	}
	db.linksLock.Lock()
	defer db.linksLock.Unlock()
	db.links[link.Id] = link
	return link, nil
}

func (db *InMemoryDB) SaveLink(link *msg.Link) error {
	db.linksLock.Lock()
	defer db.linksLock.Unlock()
	if _, ok := db.boxes[link.BoxId]; !ok {
		return errors.New("Box id doesn't exist")
	}
	db.links[link.Id] = link
	return nil
}

func (db *InMemoryDB) GetLinkById(id string) (*msg.Link, error) {
	db.linksLock.Lock()
	defer db.linksLock.Unlock()
	if link, ok := db.links[id]; ok {
		return link, nil
	}
	return nil, errors.New("No such link id")
}

func (db *InMemoryDB) GetLinks() ([]*msg.Link, error) {
	db.linksLock.Lock()
	defer db.linksLock.Unlock()
	var links []*msg.Link
	for _, v := range db.links {
		links = append(links, v)
	}
	return links, nil
}

func (db *InMemoryDB) GetLinksByBoxId(id string) ([]*msg.Link, error) {
	db.linksLock.Lock()
	defer db.linksLock.Unlock()
	var links []*msg.Link
	for _, v := range db.links {
		if v.BoxId == id {
			links = append(links, v)
		}
	}
	return links, nil
}

func (db *InMemoryDB) DeleteLink(id string) error {
	db.linksLock.Lock()
	defer db.linksLock.Unlock()
	delete(db.links, id)
	return nil
}
