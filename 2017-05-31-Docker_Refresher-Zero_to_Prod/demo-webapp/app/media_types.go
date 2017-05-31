// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "WorstSocialNetwork": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/hairyhenderson/restdemo/design
// --out=$(GOPATH)/src/github.com/hairyhenderson/restdemo
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
)

// service health (default view)
//
// Identifier: application/vnd.health+json; view=default
type Health struct {
	// the hostname
	Hostname string `form:"hostname" json:"hostname" xml:"hostname"`
	// the version
	Version string `form:"version" json:"version" xml:"version"`
}

// Validate validates the Health media type instance.
func (mt *Health) Validate() (err error) {
	if mt.Hostname == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "hostname"))
	}
	if mt.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	return
}

// A post (default view)
//
// Identifier: application/vnd.wsn.post+json; view=default
type WsnPost struct {
	// post body
	Body string `form:"body" json:"body" xml:"body"`
	// Unique post ID
	ID int `form:"id" json:"id" xml:"id"`
	// post title
	Title string `form:"title" json:"title" xml:"title"`
	// Owner's ID
	UserID int `form:"userId" json:"userId" xml:"userId"`
}

// Validate validates the WsnPost media type instance.
func (mt *WsnPost) Validate() (err error) {

	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Body == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body"))
	}

	return
}
