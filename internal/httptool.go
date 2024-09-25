package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type HTTPClient struct {
}

type Response[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type simpleHttpRequest[T any] struct {
	url     string
	method  string
	headers [][]string
	body    io.Reader
}

func NewSimpleHttpRequest[T any]() *simpleHttpRequest[T] {
	return new(simpleHttpRequest[T])
}

func (r *simpleHttpRequest[T]) SetUrl(url string) *simpleHttpRequest[T] {
	r.url = url
	return r
}

func (r *simpleHttpRequest[T]) SetMethod(method string) *simpleHttpRequest[T] {
	r.method = method
	return r
}

func (r *simpleHttpRequest[T]) SetHeaders(headers [][]string) *simpleHttpRequest[T] {
	r.headers = headers
	return r
}

func (r *simpleHttpRequest[T]) SetBody(body io.Reader) *simpleHttpRequest[T] {
	r.body = body
	return r
}

func (r *simpleHttpRequest[T]) SetJsonBody(o any) *simpleHttpRequest[T] {
	bodyBytes, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	r.body = bytes.NewReader(bodyBytes)
	return r
}

func (r *simpleHttpRequest[T]) Do() (*Response[T], error) {
	req, err := http.NewRequest(r.method, r.url, r.body)

	for _, header := range r.headers {
		if len(header) != 2 {
			panic("request Header length should be 2, in the other words, { key, value }")
		}
		req.Header.Add(header[0], header[1])
	}
	resp, err := http.DefaultClient.Do(req) // send http get request to url
	if err != nil {
		return nil, fmt.Errorf("failed to bind: %w", err)
	}
	defer resp.Body.Close() // close response body

	switch resp.StatusCode {
	case http.StatusBadRequest:
		panic(fmt.Sprintf("Wrong request param, url: %v", r.url))
	case http.StatusUnauthorized:
		return nil, errors.New("unauthorized")
	case http.StatusForbidden:
		return nil, errors.New("forbidden")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode) // return nil,error if status code is not ok
	}

	body, err := io.ReadAll(resp.Body) // read response body
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response = new(Response[T])      // create a user instance
	err = json.Unmarshal(body, response) // unmarshal json to user
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return response, nil
}

func (r *simpleHttpRequest[T]) DoBarely() (*T, error) {
	req, err := http.NewRequest(r.method, r.url, r.body)
	if err != nil {
		return nil, fmt.Errorf("svc url:%s,failed to bind: %w", r.url, err)
	}
	for _, header := range r.headers {
		if len(header) != 2 {
			panic("request Header length should be 2, in the other words, { key, value }")
		}
		req.Header.Add(header[0], header[1])
	}
	cli := http.Client{Timeout: time.Second * 10} // Set the timeout to 3 seconds
	resp, err := cli.Do(req)                      // send http get request to url

	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return nil, fmt.Errorf("svc url:%s, request timed out", r.url)
		}
		return nil, fmt.Errorf("svc url:%s, failed to bind: %w", r.url, err)
	}
	defer resp.Body.Close() // close response body

	switch resp.StatusCode {
	case http.StatusBadRequest:
		panic(fmt.Sprintf("Wrong request param, url: %v", r.url))
	case http.StatusUnauthorized:
		return nil, errors.New("unauthorized")
	case http.StatusForbidden:
		return nil, errors.New("forbidden")
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("svc url:%s,unexpected status code: %d, body: %v", r.url, resp.StatusCode, string(body)) // return nil,error if status code is not ok
	}

	body, err := io.ReadAll(resp.Body) // read response body
	if err != nil {
		return nil, fmt.Errorf("svc url:%s,failed to read response body: %w", r.url, err)
	}

	var response = new(T)                // create a user instance
	err = json.Unmarshal(body, response) // unmarshal json to user
	if err != nil {
		return nil, fmt.Errorf("svc url:%s,failed to unmarshal json: %w", r.url, err)
	}

	return response, nil
}
