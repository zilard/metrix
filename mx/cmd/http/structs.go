
package http

import (
    "net/http"
    "net/url"
)



type Client struct {
    HttpClient *http.Client
    BaseURL    *url.URL
    UserAgent  string
}

