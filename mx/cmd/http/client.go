
package http

import (
    "encoding/json"
    "net/http"
    "net/url"
    "fmt"
    "io"
    "bytes"
)


const BaseURL string = "http://localhost:8080"



func NewClient() *Client {
    c := &Client{}
    c.BaseURL, _ = url.Parse(BaseURL)
    c.HttpClient = http.DefaultClient
    c.UserAgent = fmt.Sprintf("metrix-client")
    return c
}



func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {

    rel := &url.URL{Path: path}
    u := c.BaseURL.ResolveReference(rel)

    var buf io.ReadWriter
    if body != nil {
        buf = new(bytes.Buffer)
        err := json.NewEncoder(buf).Encode(body)
	if err != nil {
            return nil, err
        }
    }

    req, err := http.NewRequest(method, u.String(), buf)
    if err != nil {
        return nil, err
    }
    if body != nil {
        req.Header.Set("Content-Type", "application/json")
    }
    req.Header.Set("Accept", "application/json")
    req.Header.Set("User-Agent", c.UserAgent)

    return req, nil

}



func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {

    resp, err := c.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    err = json.NewDecoder(resp.Body).Decode(v)

    return resp, err

}






