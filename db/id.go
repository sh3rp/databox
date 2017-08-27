package db

import (
	"math/rand"
	"strings"
	"time"

	"github.com/oklog/ulid"
)

func GenerateID() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return strings.ToLower(id.String())
}
