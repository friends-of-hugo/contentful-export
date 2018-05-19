package mapper

import (
	"encoding/json"
	"io"
)

func MapTypes(rc io.ReadCloser) (typeResult TypeResult, err error) {
	defer rc.Close()
	err = json.NewDecoder(rc).Decode(&typeResult)

	return
}

func MapItems(rc io.ReadCloser) (itemResult ItemResult, err error) {
	defer rc.Close()
	err = json.NewDecoder(rc).Decode(&itemResult)

	return
}
