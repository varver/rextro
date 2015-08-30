package rextro

import (
	"bytes"
	"github.com/Jeffail/gabs"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Tequest struct {
	Headers map[string]string
	Body    map[string]string
	UrlStr  string
}

func NewTequest(inputurl ...string) Tequest {
	startUrl := ""
	if len(inputurl) == 1 {
		startUrl = inputurl[0]
	} else if len(inputurl) > 1 {
		startUrl = strings.Join(inputurl, "/")
		startUrl = strings.Replace(startUrl, "//", "/", -1)
	}
	return Tequest{Headers: make(map[string]string), Body: make(map[string]string), UrlStr: startUrl}
}

func (req Tequest) SetUrl(inputurl ...string) Tequest {
	startUrl := ""
	if len(inputurl) == 1 {
		startUrl = inputurl[0]
	} else if len(inputurl) > 1 {
		startUrl = strings.Join(inputurl, "/")
		startUrl = strings.Replace(startUrl, "//", "/", -1)
	}
	req.UrlStr = startUrl
	return req
}

func (req Tequest) FetchJson(method string) (*gabs.Container, error) {
	back, err := req.Fetch(method)
	if err != nil {
		return nil, err
	}
	container, err := gabs.ParseJSON(back)
	return container, err
}

func (req Tequest) FetchString(method string) (string, error) {
	back, err := req.Fetch(method)
	if err != nil {
		return "", err
	}
	return string(back), err
}

func (req Tequest) Fetch(method string) ([]byte, error) {
	client := &http.Client{}

	// make body parameters
	data := url.Values{}
	if req.Body != nil || len(req.Body) > 0 {
		for bk, bv := range req.Body {
			data[bk] = []string{bv}
		}
	}

	payload := bytes.NewBufferString(data.Encode())

	// make request
	r, err := http.NewRequest(method, req.UrlStr, payload)

	// make headers
	if req.Headers != nil || len(req.Headers) > 0 {
		for hk, hv := range req.Headers {
			r.Header.Add(hk, hv)
		}
	}

	// make call
	resp, err := client.Do(r)
	if err != nil {
		return []byte(""), err
	}

	// read
	back, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return back, err
}

func Mashape(image string, mashape_key string) (*gabs.Container, error) {
	req := NewTequest("https://nuditysearch.p.mashape.com/nuditySearch/image")
	//headers
	req.Headers["X-Mashape-Key"] = mashape_key
	req.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	req.Headers["Accept"] = "application/json"
	// parameters
	req.Body["setting"] = "2"
	req.Body["objecturl"] = image
	return req.FetchJson("POST")
}
