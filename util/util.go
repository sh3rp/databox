package util

import (
	"encoding/json"
	"os"
)

func PrettyPrint(obj interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "   ")
	encoder.Encode(obj)
}
