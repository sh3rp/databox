package db

import (
	"errors"
	"sync"

	"github.com/sh3rp/databox/msg"
)

type InMemoryDB struct {
	boxes      map[msg.Key]*msg.Box
	boxesLock  *sync.Mutex
	links      map[msg.Key]*msg.Link
	linksLock  *sync.Mutex
	defaultBox *msg.Box
}

func NewInMemoryDB() BoxDB {
	db := InMemoryDB{
		boxes:     make(map[msg.Key]*msg.Box),
		links:     make(map[msg.Key]*msg.Link),
		boxesLock: new(sync.Mutex),
		linksLock: new(sync.Mutex),
	}
	//db.defaultBox, _ = db.NewBox("default", "Default box", true)
	return &db
}

func (db *InMemoryDB) NewBox(name string, description string, password []byte) (*msg.Box, error) {
	if name == "" {
		return nil, errors.New("Name must not be empty")
	}
	if description == "" {
		return nil, errors.New("Description must not be empty")
	}
	box := &msg.Box{
		Id:          NewBoxKey(),
		Name:        name,
		Description: description,
	}
	db.boxesLock.Lock()
	defer db.boxesLock.Unlock()
	db.boxes[*box.Id] = box
	db.defaultBox = db.boxes[*box.Id]
	return box, nil
}

func (db *InMemoryDB) SaveBox(box *msg.Box) error {
	db.boxesLock.Lock()
	defer db.boxesLock.Unlock()
	if box.Id == nil {
		return errors.New("No id set, cannot save box")
	}
	db.boxes[*box.Id] = box
	return nil
}

func (db *InMemoryDB) GetBoxById(id msg.Key) (*msg.Box, error) {
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

func (db *InMemoryDB) DeleteBox(id msg.Key) error {
	db.boxesLock.Lock()
	defer db.boxesLock.Unlock()
	delete(db.boxes, id)
	return nil
}

func (db *InMemoryDB) GetDefaultBox() (*msg.Box, error) {
	return db.defaultBox, nil
}

func (db *InMemoryDB) NewLink(name string, url string, boxId msg.Key) (*msg.Link, error) {
	if boxId.Id == "" {
		return nil, errors.New("Cannot specify empty box ID")
	}

	_, err := db.GetBoxById(boxId)

	if err != nil {
		return nil, err
	}

	link := &msg.Link{
		Id:   NewLinkKey(&boxId),
		Name: name,
		Url:  url,
	}
	db.linksLock.Lock()
	defer db.linksLock.Unlock()
	db.links[*link.Id] = link
	return link, nil
}

func (db *InMemoryDB) SaveLink(link *msg.Link) error {
	db.linksLock.Lock()
	defer db.linksLock.Unlock()

	box, err := db.GetBoxById(msg.Key{
		Type: msg.Key_BOX,
		Id:   link.Id.BoxId,
	})

	if err != nil {
		return err
	}

	if _, ok := db.boxes[*box.Id]; !ok {
		return errors.New("Box id doesn't exist")
	}
	db.links[*link.Id] = link
	return nil
}

func (db *InMemoryDB) GetLinkById(id msg.Key) (*msg.Link, error) {
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

func (db *InMemoryDB) GetLinksByBoxId(id msg.Key) ([]*msg.Link, error) {
	if _, ok := db.boxes[id]; !ok {
		return nil, errors.New("No such box id")
	}
	db.linksLock.Lock()
	defer db.linksLock.Unlock()
	var links []*msg.Link
	for _, v := range db.links {
		if v.Id.BoxId == id.Id {
			links = append(links, v)
		}
	}
	return links, nil
}

func (db *InMemoryDB) DeleteLink(id msg.Key) error {

	_, err := db.GetBoxById(msg.Key{
		Type: msg.Key_BOX,
		Id:   id.BoxId,
	})

	if err != nil {
		return err
	}

	db.linksLock.Lock()
	defer db.linksLock.Unlock()
	delete(db.links, id)
	return nil
}
