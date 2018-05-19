package read

import "errors"

type Reader struct {
	Store Store
}

func (r *Reader) ViewFromFile(fileName string) (body string, err error) {
	result, err := r.Store.ReadFromFile(fileName)
	if result == nil && err == nil {
		err = errors.New("File is Empty")
	}

	return string(result), err
}
