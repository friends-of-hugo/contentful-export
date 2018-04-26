package read

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type HttpGetter struct {
}

// Get makes an http get request and throws an Error if the response
// statuscode is not 200
func (hg HttpGetter) Get(url string) (result io.ReadCloser, err error) {
	resp, err := myClient.Get(url)
	if resp.StatusCode != 200 && err == nil {
		err = fmt.Errorf("Http request failed: %s", resp.Status)
	}
	return resp.Body, err
}
