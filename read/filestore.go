package read

import (
	"io/ioutil"
)

type FileStore struct{}

func (fs FileStore) ReadFromFile(filename string) (body []byte, err error) {
	bytes, err := ioutil.ReadFile(filename)

	return bytes, err
}
