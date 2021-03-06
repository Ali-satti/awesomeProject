// Code generated by go-swagger; DO NOT EDIT.

package read_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListItemParams creates a new ListItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListItemParams() *ListItemParams {
	return &ListItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListItemParamsWithTimeout creates a new ListItemParams object
// with the ability to set a timeout on a request.
func NewListItemParamsWithTimeout(timeout time.Duration) *ListItemParams {
	return &ListItemParams{
		timeout: timeout,
	}
}

// NewListItemParamsWithContext creates a new ListItemParams object
// with the ability to set a context for a request.
func NewListItemParamsWithContext(ctx context.Context) *ListItemParams {
	return &ListItemParams{
		Context: ctx,
	}
}

// NewListItemParamsWithHTTPClient creates a new ListItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewListItemParamsWithHTTPClient(client *http.Client) *ListItemParams {
	return &ListItemParams{
		HTTPClient: client,
	}
}

/* ListItemParams contains all the parameters to send to the API endpoint
   for the list item operation.

   Typically these are written to a http.Request.
*/
type ListItemParams struct {

	// Kind.
	Kind *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListItemParams) WithDefaults() *ListItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListItemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list item params
func (o *ListItemParams) WithTimeout(timeout time.Duration) *ListItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list item params
func (o *ListItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list item params
func (o *ListItemParams) WithContext(ctx context.Context) *ListItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list item params
func (o *ListItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list item params
func (o *ListItemParams) WithHTTPClient(client *http.Client) *ListItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list item params
func (o *ListItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKind adds the kind to the list item params
func (o *ListItemParams) WithKind(kind *string) *ListItemParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the list item params
func (o *ListItemParams) SetKind(kind *string) {
	o.Kind = kind
}

// WriteToRequest writes these params to a swagger request
func (o *ListItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Kind != nil {

		// query param kind
		var qrKind string

		if o.Kind != nil {
			qrKind = *o.Kind
		}
		qKind := qrKind
		if qKind != "" {

			if err := r.SetQueryParam("kind", qKind); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
