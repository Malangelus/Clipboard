// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Clipboard": actions Resource Client
//
// Command:
// $ goagen
// --design=sebs.software\Malangelus\Clipboard\design
// --out=$(GOPATH)\src\sebs.software\Malangelus\Clipboard
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	uuid "github.com/goadesign/goa/uuid"
	"net/http"
	"net/url"
)

// RetrieveActionsPath computes a request path to the retrieve action of actions.
func RetrieveActionsPath(guid uuid.UUID) string {
	param0 := guid.String()

	return fmt.Sprintf("/retrieve/%s", param0)
}

// RetrieveActionsPath2 computes a request path to the retrieve action of actions.
func RetrieveActionsPath2() string {

	return fmt.Sprintf("/retrieve")
}

// Stores a string value in memory
func (c *Client) RetrieveActions(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewRetrieveActionsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRetrieveActionsRequest create the request corresponding to the retrieve action endpoint of the actions resource.
func (c *Client) NewRetrieveActionsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// StoreActionsPath computes a request path to the store action of actions.
func StoreActionsPath(value string) string {
	param0 := value

	return fmt.Sprintf("/store/%s", param0)
}

// StoreActionsPath2 computes a request path to the store action of actions.
func StoreActionsPath2() string {

	return fmt.Sprintf("/store")
}

// Stores a string value in memory
func (c *Client) StoreActions(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewStoreActionsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewStoreActionsRequest create the request corresponding to the store action endpoint of the actions resource.
func (c *Client) NewStoreActionsRequest(ctx context.Context, path string) (*http.Request, error) {
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