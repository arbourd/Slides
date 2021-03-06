// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "WorstSocialNetwork": health Resource Client
//
// Command:
// $ goagen
// --design=github.com/hairyhenderson/restdemo/design
// --out=$(GOPATH)/src/github.com/hairyhenderson/restdemo
// --tool=restdemo-cli
// --version=v1.2.0-dirty

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ShowHealthPath computes a request path to the show action of health.
func ShowHealthPath() string {

	return fmt.Sprintf("/health")
}

// show health
func (c *Client) ShowHealth(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowHealthRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowHealthRequest create the request corresponding to the show action endpoint of the health resource.
func (c *Client) NewShowHealthRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
