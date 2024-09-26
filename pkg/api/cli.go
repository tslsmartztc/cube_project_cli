package api

import (
	"fmt"
)

type Client struct {
	Host    string
	Port    int
	BaseUrl string
}

type ClientOpt struct {
	BaseUrl string
}

func NewClient(host string, port int, opts ...ClientOpt) *Client {
	if len(opts) > 1 {
		panic("too many options")
	}

	c := &Client{
		Host: host,
		Port: port,
	}

	if len(opts) == 1 {
		c.BaseUrl = opts[0].BaseUrl
	}
	return c
}

func (c *Client) getClient() string {
	return c.appendUrl(fmt.Sprintf("%s:%d", c.Host, c.Port), c.BaseUrl)
}

func (c *Client) httpUrl(path ...string) string {
	return c.appendUrl(fmt.Sprintf("http://%s", c.getClient()), path...)
}

func (c *Client) appendUrlParams(url string, params ...string) string {
	if len(params) == 0 {
		return url
	}

	if url[len(url)-1] != '?' {
		url += "?"
	}
	for i := 0; i < len(params); i += 2 {
		if i/2 != 0 {
			url += "&"
		}
		if i+1 >= len(params) {
			break
		}
		url += fmt.Sprintf("%s=%s", params[i], params[i+1])
	}
	return url
}

func (c *Client) appendUrl(url string, path ...string) string {
	appends := ""
	for _, p := range path {
		if p[0] != '/' {
			p = "/" + p
		}
		if p[len(p)-1] == '/' {
			p = p[:len(p)-1]
		}
		appends += p
	}
	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}
	return url + appends
}
