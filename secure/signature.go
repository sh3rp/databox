package secure

import (
	"crypto/sha256"

	"github.com/sh3rp/databox/msg"
)

func GetSignature(password []byte, box *msg.Box) []byte {
	hasher := sha256.New()
	hasher.Write(password)
	hasher.Write([]byte(box.Id.Id))
	return hasher.Sum(nil)
}
