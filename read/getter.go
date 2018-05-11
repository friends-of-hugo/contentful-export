package read

import (
	"io"
)

type Getter interface {
	Get(url string) (result io.ReadCloser, err error)
}
