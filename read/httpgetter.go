package read

import (
	"io"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type HttpGetter struct {
}

func (hg HttpGetter) Get(url string) (result io.ReadCloser, err error) {
	resp, err := myClient.Get(url)
	return resp.Body, err
}
