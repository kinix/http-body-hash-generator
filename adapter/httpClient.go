package adapter

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// If you want to use https, change this
const defaultProtocol = "http"

type HttpClient struct {
}

// Send a HTTP GET request and return the response body
func (c *HttpClient) GetBody(url string) (string, error) {
	// If the url doesn't begin with one of supported protocol, prepend default protocol
	if len(url) < 8 || (url[:7] != "http://" && url[:8] != "https://") {
		url = fmt.Sprintf("%s://%s", defaultProtocol, url)
	}

	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// Read body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
