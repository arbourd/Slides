// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "WorstSocialNetwork": post Resource Client
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
	"strconv"
)

// ShowPostPath computes a request path to the show action of post.
func ShowPostPath(postID int) string {
	param0 := strconv.Itoa(postID)

	return fmt.Sprintf("/posts/%s", param0)
}

// Get post by id
func (c *Client) ShowPost(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowPostRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowPostRequest create the request corresponding to the show action endpoint of the post resource.
func (c *Client) NewShowPostRequest(ctx context.Context, path string) (*http.Request, error) {
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
