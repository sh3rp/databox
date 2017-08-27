package db

import (
	"errors"
	"sync"

	"github.com/sh3rp/databox/msg"
)

type InMemoryDB struct {
	boxes     map[int64]*msg.Box
	boxesLock *sync.Mutex
	boxId     int64
	links     map[int64]*msg.Link
	linksLock *sync.Mutex
	linkId    int64
}

func NewInMemoryDB() *InMemoryDB {
	db := InMemoryDB{
		boxes:     make(map[int64]*msg.Box),
		links:     make(map[int64]*msg.Link),
		boxesLock: new(sync.Mutex),
		linksLock: new(sync.Mutex),
	}
	db.NewBox("default")
	return &db
}

func (db *InMemoryDB) NewBox(name string) (*msg.Box, error) {
	box := &msg.Box{
		Id:   db.newBoxId(),
		Name: name,
	}
	db.boxesLock.Lock()
	defer db.boxesLock.Unlock()
	db.boxes[box.Id] = box
	return box, nil
}

func (db *InMemoryDB) SaveBox(box *msg.Box) error {
	db.boxesLock.Lock()
	defer db.boxesLock.Unlock()
	if box.Id == 0 {
		box.Id = db.newBoxId()
	}
	db.boxes[box.Id] = box
	return nil
}

func (db *InMemoryDB) GetBoxById(id int64) (*msg.Box, error) {
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

func (db *InMemoryDB) DeleteBox(id int64) error {
	db.boxesLock.Lock()
	defer db.boxesLock.Unlock()
	delete(db.boxes, id)
	return nil
}

func (db *InMemoryDB) NewLink(name string, url string, boxId int64) (*msg.Link, error) {
	link := &msg.Link{
		Id:    db.newLinkId(),
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

func (db *InMemoryDB) GetLinkById(id int64) (*msg.Link, error) {
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

func (db *InMemoryDB) GetLinksByBoxId(id int64) ([]*msg.Link, error) {
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

func (db *InMemoryDB) DeleteLink(id int64) error {
	db.linksLock.Lock()
	defer db.linksLock.Unlock()
	delete(db.links, id)
	return nil
}

func (db *InMemoryDB) newBoxId() int64 {
	db.boxId = db.boxId + 1
	return db.boxId
}

func (db *InMemoryDB) newLinkId() int64 {
	db.linkId = db.linkId + 1
	return db.linkId
}
