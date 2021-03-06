// Code generated by go-swagger; DO NOT EDIT.

package postgre_sql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new postgre sql API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for postgre sql API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddMixin4 add mixin4 API
*/
func (a *Client) AddMixin4(params *AddMixin4Params) (*AddMixin4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddMixin4Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddMixin4",
		Method:             "POST",
		PathPattern:        "/v0/postgresql",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddMixin4Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddMixin4OK), nil

}

/*
ListMixin4 list mixin4 API
*/
func (a *Client) ListMixin4(params *ListMixin4Params) (*ListMixin4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListMixin4Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListMixin4",
		Method:             "GET",
		PathPattern:        "/v0/postgresql",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListMixin4Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListMixin4OK), nil

}

/*
RemoveMixin4 remove mixin4 API
*/
func (a *Client) RemoveMixin4(params *RemoveMixin4Params) (*RemoveMixin4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveMixin4Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RemoveMixin4",
		Method:             "DELETE",
		PathPattern:        "/v0/postgresql/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveMixin4Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveMixin4OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
