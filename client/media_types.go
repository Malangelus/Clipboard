// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Clipboard": Application Media Types
//
// Command:
// $ goagen
// --design=sebs.software\Malangelus\Clipboard\design
// --out=$(GOPATH)\src\sebs.software\Malangelus\Clipboard
// --version=v1.3.1

package client

import (
	"net/http"
)

// ClipboardRetrieve media type (default view)
//
// Identifier: application/vnd.clipboard.retrieve; view=default
type ClipboardRetrieve struct {
	Value *string `form:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty" xml:"value,omitempty"`
}

// DecodeClipboardRetrieve decodes the ClipboardRetrieve instance encoded in resp body.
func (c *Client) DecodeClipboardRetrieve(resp *http.Response) (*ClipboardRetrieve, error) {
	var decoded ClipboardRetrieve
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ClipboardStore media type (default view)
//
// Identifier: application/vnd.clipboard.store; view=default
type ClipboardStore struct {
	GUID *string `form:"guid,omitempty" json:"guid,omitempty" yaml:"guid,omitempty" xml:"guid,omitempty"`
}

// DecodeClipboardStore decodes the ClipboardStore instance encoded in resp body.
func (c *Client) DecodeClipboardStore(resp *http.Response) (*ClipboardStore, error) {
	var decoded ClipboardStore
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
