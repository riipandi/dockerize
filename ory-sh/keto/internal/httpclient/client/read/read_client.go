// Code generated by go-swagger; DO NOT EDIT.

package read

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new read API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for read API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCheck(params *GetCheckParams, opts ...ClientOption) (*GetCheckOK, error)

	GetExpand(params *GetExpandParams, opts ...ClientOption) (*GetExpandOK, error)

	GetRelationTuples(params *GetRelationTuplesParams, opts ...ClientOption) (*GetRelationTuplesOK, error)

	PostCheck(params *PostCheckParams, opts ...ClientOption) (*PostCheckOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCheck checks a relation tuple

  To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
*/
func (a *Client) GetCheck(params *GetCheckParams, opts ...ClientOption) (*GetCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCheck",
		Method:             "GET",
		PathPattern:        "/relation-tuples/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCheckReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetExpand expands a relation tuple

  Use this endpoint to expand a relation tuple.
*/
func (a *Client) GetExpand(params *GetExpandParams, opts ...ClientOption) (*GetExpandOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExpandParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExpand",
		Method:             "GET",
		PathPattern:        "/relation-tuples/expand",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExpandReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetExpandOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExpand: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRelationTuples queries relation tuples

  Get all relation tuples that match the query. Only the namespace field is required.
*/
func (a *Client) GetRelationTuples(params *GetRelationTuplesParams, opts ...ClientOption) (*GetRelationTuplesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRelationTuplesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRelationTuples",
		Method:             "GET",
		PathPattern:        "/relation-tuples",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRelationTuplesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRelationTuplesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRelationTuples: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostCheck checks a relation tuple

  To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
*/
func (a *Client) PostCheck(params *PostCheckParams, opts ...ClientOption) (*PostCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postCheck",
		Method:             "POST",
		PathPattern:        "/relation-tuples/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostCheckReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
