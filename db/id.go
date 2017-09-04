package db

import (
	"math/rand"
	"strings"
	"time"

	"github.com/oklog/ulid"
	"github.com/sh3rp/databox/msg"
)

func NewBoxKey() *msg.Key {
	return &msg.Key{
		Type: msg.Key_BOX,
		Id:   GenerateID(),
	}
}

func NewLinkKey(boxKey *msg.Key) *msg.Key {
	return &msg.Key{
		Type:  msg.Key_LINK,
		Id:    GenerateID(),
		BoxId: boxKey.Id,
	}
}

func GenerateID() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return strings.ToLower(id.String())
}
