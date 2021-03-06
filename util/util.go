package util

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"

	"github.com/sh3rp/databox/msg"
)

var V_MAJOR = 0
var V_MINOR = 1
var V_PATCH = 0

func PrettyPrint(obj interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "   ")
	encoder.Encode(obj)
}

// laziness, will move this later
func GetVersion() *msg.Version {
	return &msg.Version{
		Major: int32(V_MAJOR),
		Minor: int32(V_MINOR),
		Patch: int32(V_PATCH),
	}
}

func GetPassHash(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
