package lib

import (
	"encoding/json"
	"os"
)

func ParseJSON(path string, decoder ...interface{}) (err error) {
	targetFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer targetFile.Close()

	err = json.NewDecoder(targetFile).Decode(decoder[0])
	return
}
